package main

import (
	"ModTask/internal/db"
	"ModTask/internal/handlers"
	"ModTask/internal/taskService"
	"ModTask/internal/userService"
	"ModTask/internal/web/tasks"
	"ModTask/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not to database: %v", err)
	}

	tasksRepo := taskService.NewTaskRepository(database)
	tasksService := taskService.NewTaskService(tasksRepo)
	tasksHandlers := handlers.NewTaskHandlers(tasksService)

	userRepo := userService.NewUserRepository(database)
	userService := userService.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandlers(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(tasksHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	strictUserHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("filed to start with err :%v", err)
	}
}
