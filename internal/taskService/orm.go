package taskService

type Task struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Task   string `json:"task"`
	UserID int    `gorm:"not null" json:"user_id"`
}

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []Task `gorm:"foreignKey:UserID; references:ID" json:"tasks"`
}
