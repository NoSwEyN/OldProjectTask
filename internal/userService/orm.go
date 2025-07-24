package userService

type User struct {
	ID       int    `gorm:"primaryKey;autoMIncrement" json:"id"`
	Email    string `lson:"email"`
	Password string `json:"password"`
}
