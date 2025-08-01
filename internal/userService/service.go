package userService

type UserService interface {
	PostService(email, password string) (User, error)
	GetAllService() ([]User, error)
	GetServiceByID(id int) (User, error)
	UpdateService(id int, email, password string) (User, error)
	DeleteService(id int) error
	GetAllUsersIdService(userID int) (*User, error)
}

type usersService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &usersService{repo: r}
}

func (s *usersService) PostService(email, password string) (User, error) {
	newUser := User{Email: email, Password: password}
	return s.repo.PostRepository(newUser)
}

func (s *usersService) GetAllService() ([]User, error) {
	return s.repo.GetAllRepository()
}

func (s *usersService) GetServiceByID(id int) (User, error) {
	return s.repo.GetRepositoryID(id)
}

func (s *usersService) UpdateService(id int, email, password string) (User, error) {
	users, err := s.repo.GetRepositoryID(id)
	if err != nil {
		return User{}, err
	}

	if email != "" {
		users.Email = email
	}

	if password != "" {
		users.Password = password
	}

	update, err := s.repo.UpdateRepository(users)
	if err != nil {
		return User{}, err
	}
	return update, nil
}

func (s *usersService) DeleteService(id int) error {
	return s.repo.DeleteRepository(id)
}

func (s *usersService) GetAllUsersIdService(userID int) (*User, error) {
	user, err := s.repo.GetTasksForUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil

}
