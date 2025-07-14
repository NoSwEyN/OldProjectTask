package taskService

type TaskService interface {
	PostService(task string) (Task, error)
	GetAllService() ([]Task, error)
	GetServiceID(id int) (Task, error)
	UpdateService(id int, task Task) (Task, error)
	DeleteService(id int) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) PostService(task string) (Task, error) {
	newTask := Task{Task: task}
	return s.repo.PostRepository(newTask)
}

func (s *taskService) GetAllService() ([]Task, error) {
	return s.repo.GetAllRepository()
}

func (s *taskService) GetServiceID(id int) (Task, error) {
	return s.repo.GetRepositoryID(id)
}

func (s *taskService) UpdateService(id int, task Task) (Task, error) {
	tasks, err := s.repo.GetRepositoryID(id)
	if err != nil {
		return Task{}, err
	}

	tasks.Task = task.Task

	update, err := s.repo.UpdateRepository(tasks)
	if err != nil {
		return Task{}, err
	}
	return update, nil
}

func (s *taskService) DeleteService(id int) error {
	return s.repo.DeleteRepository(id)
}
