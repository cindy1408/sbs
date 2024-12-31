package graph

//type Todo struct {
//	ID     string `gorm:"primaryKey"`
//	Text   string `json:"text"`
//	Done   bool   `json:"done"`
//	UserID string `json:"userId" gorm:"type:varchar(255);not null;"`                  // Foreign key to User
//	User   *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Define relationship to User
//}
//
//type User struct {
//	ID    string `gorm:"primaryKey"`
//	Name  string `json:"name"`
//	Todos []Todo `json:"todos" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relation
//}
