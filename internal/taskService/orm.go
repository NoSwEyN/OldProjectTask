package taskService

type Task struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Task string `josn:"task"`
}
