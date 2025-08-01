package userService

import "ModTask/internal/taskService"

type User struct {
	ID       int                 `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string              `json:"email"`
	Password string              `json:"password"`
	Tasks    []*taskService.Task `gorm:"foreignKey:UserID; references:ID" json:"tasks"`
}
