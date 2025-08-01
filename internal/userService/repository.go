package userService

import (
	"ModTask/internal/taskService"

	"gorm.io/gorm"
)

type UserRepository interface {
	PostRepository(user taskService.User) (taskService.User, error)
	GetAllRepository() ([]taskService.User, error)
	GetRepositoryID(id int) (taskService.User, error)
	UpdateRepository(user taskService.User) (taskService.User, error)
	DeleteRepository(id int) error
	GetTasksForUser(userID int) (*taskService.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) PostRepository(user taskService.User) (taskService.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetAllRepository() ([]taskService.User, error) {
	var user []taskService.User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *userRepository) GetRepositoryID(id int) (taskService.User, error) {
	var user taskService.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateRepository(user taskService.User) (taskService.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteRepository(id int) error {
	return r.db.Delete(&taskService.User{}, "id = ?", id).Error
}

func (r *userRepository) GetTasksForUser(userID int) (*taskService.User, error) {
	var user taskService.User
	err := r.db.Preload("Tasks").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
