package handlers

import (
	"ModTask/internal/taskService"
	"ModTask/internal/web/tasks"
	"context"
	"fmt"
)

type TaskHandlers struct {
	service taskService.TaskService
}

func NewTaskHandlers(s taskService.TaskService) *TaskHandlers {
	return &TaskHandlers{service: s}
}

func (h *TaskHandlers) GetTasks(_ context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllService()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandlers) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	if taskRequest == nil || taskRequest.Task == nil || taskRequest.UserId == nil {
		return nil, fmt.Errorf("task and user_id are required")
	}

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.service.PostService(taskToCreate.Task, *taskRequest.UserId)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		UserId: &createdTask.UserID,
	}
	return response, nil
}

func (h *TaskHandlers) PutTasksId(_ context.Context, request tasks.PutTasksIdRequestObject) (tasks.PutTasksIdResponseObject, error) {
	id := request.Id

	var err error

	body := request.Body

	if body.Task == nil {
		return nil, err
	}

	updatedTask := taskService.Task{
		ID:   id,
		Task: *body.Task,
	}

	update, err := h.service.UpdateService(id, updatedTask)
	if err != nil {
		return nil, err
	}

	res := tasks.PutTasksId200JSONResponse{
		Id:   &id,
		Task: &update.Task,
	}
	return res, nil
}

func (h *TaskHandlers) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	if err := h.service.DeleteService(id); err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204JSONResponse{}, nil
}
