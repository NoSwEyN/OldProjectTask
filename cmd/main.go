package main

import (
	"ModTask/internal/db"
	"ModTask/internal/handlers"
	"ModTask/internal/taskService"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not to database: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandlers(taskService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/tasks", taskHandler.CreateHandler)
	e.GET("/tasks", taskHandler.GetHandler)
	e.PUT("/tasks/:id", taskHandler.UpdateHandler)
	e.DELETE("/tasks/:id", taskHandler.DeleteHandler)

	e.Start("localhost:8080")
}
