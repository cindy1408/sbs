package graph

type Todo struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId" gorm:"type:varchar(255);not null;"`   // Foreign key to User
	User   *User  `json:"user" gorm:"foreignKey:UserID;references:ID"` // Relationship to User, uses UserID as the foreign key
}

type User struct {
	ID    string  `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Todos []*Todo `json:"todos,omitempty" gorm:"foreignKey:UserID"` // Explicitly defining foreign key in the relationship
}
