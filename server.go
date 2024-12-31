package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"stepbystep.com/m/graph"
	"stepbystep.com/m/graph/model"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to default postgres database first
	defaultDB, err := gorm.Open(postgres.Open("host=localhost user=cindycheung dbname=postgres port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to default database:", err)
	}

	// Create todo database if it doesn't exist
	var exists bool
	defaultDB.Raw("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = 'todo')").Scan(&exists)
	if !exists {
		log.Println("Creating todo database...")
		err := defaultDB.Exec("CREATE DATABASE todo").Error
		if err != nil {
			log.Fatal("Failed to create todo database:", err)
		}
		log.Println("Todo database created successfully")
	}

	// Connect to todo database
	db, err := gorm.Open(postgres.Open("host=localhost user=cindycheung dbname=todo port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to todo database:", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&model.Todo{},
		&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

	resolver := &graph.Resolver{Db: db}
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
