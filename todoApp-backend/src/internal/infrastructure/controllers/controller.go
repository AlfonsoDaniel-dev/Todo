package controllers

import (
	"github.com/labstack/echo/v4"
	"todoApp-backend/src/internal/app"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/middlewares"
)

type handler struct {
	UserServices app.UserServices
	TaskServices app.TaskServices
}

func newHandler(userRepository domain.UserRepository, taskRepository domain.TaskRepository) *handler {

	userService := app.NewUserServices(userRepository)
	taskService := app.NewTaskServices(taskRepository)

	return &handler{
		UserServices: *userService,
		TaskServices: *taskService,
	}
}

type controller struct {
	E              *echo.Echo
	handlers       *handler
	UserRepository domain.UserRepository
	TaskRepository domain.TaskRepository
}

func NewController(userRepository domain.UserRepository, taskRepository domain.TaskRepository) *controller {

	handler := newHandler(userRepository, taskRepository)

	return &controller{
		E:              echo.New(),
		handlers:       handler,
		UserRepository: userRepository,
		TaskRepository: taskRepository,
	}
}

func (C *controller) MountEndpoints() {
	C.UserRoutes()
}

func (C *controller) UserRoutes() {

	userPublicRoutes := C.E.Group("/user")
	userPublicRoutes.Use(middlewares.LogRequest)

	userPublicRoutes.GET("/createuser", C.handlers.createUser)

}
