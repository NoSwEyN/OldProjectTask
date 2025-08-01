package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	PostRepository(user User) (User, error)
	GetAllRepository() ([]User, error)
	GetRepositoryID(id int) (User, error)
	UpdateRepository(user User) (User, error)
	DeleteRepository(id int) error
	GetTasksForUser(userID int) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) PostRepository(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetAllRepository() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *userRepository) GetRepositoryID(id int) (User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateRepository(user User) (User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteRepository(id int) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}

func (r *userRepository) GetTasksForUser(userID int) (*User, error) {
	var user User
	err := r.db.Preload("Tasks").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
