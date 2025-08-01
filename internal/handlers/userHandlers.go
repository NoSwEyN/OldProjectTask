package handlers

import (
	"ModTask/internal/userService"
	"ModTask/internal/web/users"
	"context"
	"fmt"
)

type UserHandlers struct {
	service userService.UserService
}

func NewUserHandlers(s userService.UserService) *UserHandlers {
	return &UserHandlers{service: s}
}

func (h *UserHandlers) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {

	usersRequest := request.Body
	if usersRequest == nil {
		return nil, fmt.Errorf("request body is required")
	}

	if usersRequest.Email == nil {
		return nil, fmt.Errorf("email is required")
	}

	if usersRequest.Password == nil {
		return nil, fmt.Errorf("password is required")
	}

	userToCreate := userService.User{
		Email:    *usersRequest.Email,
		Password: *usersRequest.Password,
	}

	createdUser, err := h.service.PostService(userToCreate.Email, userToCreate.Password)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (h *UserHandlers) GetUsers(_ context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllService()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandlers) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	body := request.Body

	existingUser, err := h.service.GetServiceByID(id)
	if err != nil {
		return nil, err
	}

	if body.Email != nil {
		existingUser.Email = *body.Email
	}
	if body.Password != nil {
		existingUser.Password = *body.Password
	}

	updated, err := h.service.UpdateService(id, existingUser.Email, existingUser.Password)
	if err != nil {
		return nil, err
	}

	res := users.PatchUsersId200JSONResponse{
		Id:       &updated.ID,
		Email:    &updated.Email,
		Password: &updated.Password,
	}
	return res, nil
}

func (h *UserHandlers) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	if err := h.service.DeleteService(id); err != nil {
		return nil, err
	}
	return users.DeleteUsersId204JSONResponse{}, nil
}

func (h *UserHandlers) GetUsersIdTasks(_ context.Context, request users.GetUsersIdTasksRequestObject) (users.GetUsersIdTasksResponseObject, error) {
	id := request.Id

	user, err := h.service.GetAllUsersIdService(id)
	if err != nil {
		return nil, err
	}

	var tasks []users.Task
	for _, t := range user.Tasks {
		id := t.ID
		taskName := t.Task
		userID := t.UserID
		tasks = append(tasks, users.Task{
			Id:     &id,
			Task:   &taskName,
			UserId: &userID,
		})
	}

	return users.GetUsersIdTasks200JSONResponse(tasks), nil
}
