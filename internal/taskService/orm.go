package taskService

type Task struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Task   string `json:"task"`
	UserID int    `gorm:"not null" json:"user_id"`
}
