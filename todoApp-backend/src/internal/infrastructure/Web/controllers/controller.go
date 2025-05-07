package controllers

import "C"
import (
	"github.com/labstack/echo/v4"
	"io"
	"os"
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

	title := os.Getenv("APP_PAGE_TITLE")
	if title == "" {
		panic("APP_PAGE_TITLE isn't set")
	}

	return &handler{
		UserServices: *userService,
		TaskServices: *taskService,
	}
}

type controller struct {
	E              *echo.Echo
	handlers       *handler
	AppRoutes      *echo.Group
	ApiGroups      *echo.Group
	UserRepository domain.UserRepository
	TaskRepository domain.TaskRepository
}

type templateRenderer interface {
	Render(w io.Writer, name string, data interface{}, c echo.Context) error
}

func NewController(userRepository domain.UserRepository, taskRepository domain.TaskRepository, echo *echo.Echo, renderer templateRenderer) *controller {

	handler := newHandler(userRepository, taskRepository)

	apiGroups := echo.Group("/api/v1")

	appRoutes := apiGroups.Group("/")

	echo.Renderer = renderer

	return &controller{
		E:              echo,
		handlers:       handler,
		ApiGroups:      apiGroups,
		AppRoutes:      appRoutes,
		UserRepository: userRepository,
		TaskRepository: taskRepository,
	}
}

func (C *controller) MountEndpoints() {

	C.AppRoutes.GET("/home", C.handlers.HomePage)

	C.UserRoutes()
}

func (C *controller) UserRoutes() {
	userPublicRoutes := C.ApiGroups.Group("/user")
	userPublicRoutes.Use(middlewares2.LogRequest)

	userPublicRoutes.GET("/create", C.handlers.CreateUser)
	userPublicRoutes.POST("/login", C.handlers.Login)
	userPublicRoutes.POST("/login/google/:code", C.handlers.LoginOauth)

	userPrivateRoutes := C.ApiGroups.Group("/user/private")

	userPrivateRoutes.Use(middlewares2.AuthMiddleWare)

}
