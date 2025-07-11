package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	PostRepository(task Task) (Task, error)
	GetAllRepository() ([]Task, error)
	GetRepositoryID(id int) (Task, error)
	UpdateRepository(task Task, update string) (Task, error)
	DeleteRepository(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) PostRepository(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) GetAllRepository() ([]Task, error) {
	var task []Task
	err := r.db.Find(&task).Error
	return task, err
}

func (r *taskRepository) GetRepositoryID(id int) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) UpdateRepository(task Task, update string) (Task, error) {
	task.Task = update
	err := r.db.Save(&task).Error
	return task, err
}

func (r *taskRepository) DeleteRepository(id int) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
