package controllers

import "C"
import (
	"github.com/labstack/echo/v4"
	"todoApp-backend/src/internal/app/task"
	"todoApp-backend/src/internal/app/user"
	"todoApp-backend/src/internal/domain"
	middlewares2 "todoApp-backend/src/internal/infrastructure/Web/middlewares"
)

type handler struct {
	UserServices user.UserServices
	TaskServices task.TaskServices
}

func newHandler(userRepository domain.UserRepository, taskRepository domain.TaskRepository) *handler {

	userService := user.NewUserServices(userRepository)
	taskService := task.NewTaskServices(taskRepository)

	return &handler{
		UserServices: *userService,
		TaskServices: *taskService,
	}
}

type controller struct {
	E              *echo.Echo
	handlers       *handler
	Groups         *echo.Group
	UserRepository domain.UserRepository
	TaskRepository domain.TaskRepository
}

func NewController(userRepository domain.UserRepository, taskRepository domain.TaskRepository, echo *echo.Echo) *controller {

	handler := newHandler(userRepository, taskRepository)

	Groups := echo.Group("/api/v1")

	return &controller{
		E:              echo,
		handlers:       handler,
		Groups:         Groups,
		UserRepository: userRepository,
		TaskRepository: taskRepository,
	}
}

func (C *controller) MountEndpoints() {
	C.UserRoutes()
}

func (C *controller) UserRoutes() {
	userPublicRoutes := C.Groups.Group("/user")
	userPublicRoutes.Use(middlewares2.LogRequest)

	userPublicRoutes.GET("/create", C.handlers.CreateUser)
	userPublicRoutes.POST("/login", C.handlers.Login)
	userPublicRoutes.POST("/login/google/:code", C.handlers.LoginOauth)

	userPrivateRoutes := C.Groups.Group("/user/private")

	userPrivateRoutes.Use(middlewares2.AuthMiddleWare)

}
