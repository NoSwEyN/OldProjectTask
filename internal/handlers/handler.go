package handlers

import (
	"ModTask/internal/taskService"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TaskHandlers struct {
	service taskService.TaskService
}

func NewTaskHandlers(s taskService.TaskService) *TaskHandlers {
	return &TaskHandlers{service: s}
}

func (h *TaskHandlers) CreateHandler(c echo.Context) error {
	var task taskService.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request"})
	}

	create, err := h.service.PostService(task.Task)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not Found"})
	}
	return c.JSON(http.StatusCreated, create)
}

func (h *TaskHandlers) GetHandler(c echo.Context) error {
	task, err := h.service.GetAllService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Server ERROR"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandlers) UpdateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var task taskService.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request"})
	}

	update, err := h.service.UpdateService(id, task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Server ERROR"})
	}
	return c.JSON(http.StatusOK, update)
}

func (h *TaskHandlers) DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request"})
	}
	err = h.service.DeleteService(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Server ERROR"})
	}
	return c.NoContent(http.StatusNoContent)
}
