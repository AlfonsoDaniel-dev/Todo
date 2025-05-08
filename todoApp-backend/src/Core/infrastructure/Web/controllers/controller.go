package controllers

import "C"
import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"os"
	"todoApp-backend/src/Core/app/task"
	"todoApp-backend/src/Core/app/user"
	"todoApp-backend/src/Core/domain"
	middlewares2 "todoApp-backend/src/Core/infrastructure/Web/middlewares"
	"todoApp-backend/src/Core/infrastructure/repositories/data"
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

func NewController(db *sql.DB, echo *echo.Echo) *controller {

	userRepository := data.NewUserRepository(db)

	taskRepository := data.NewTaskRepository(db)

	handler := newHandler(userRepository, taskRepository)

	apiGroups := echo.Group("/api/v1")

	appRoutes := apiGroups.Group("/")

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

	C.E.Static("/static", "../static/pages/static")

	C.AppRoutes.GET("/", C.handlers.HomePage)
	C.AppRoutes.GET("/login", C.handlers.LoginPage)
	C.AppRoutes.GET("/signup", C.handlers.SignUpPage)
	C.AppRoutes.GET("/faq", C.handlers.FaqPage)

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

func (C *controller) TaskRoutes() {
	taskPublicRoutes := C.ApiGroups.Group("/task")
	taskPublicRoutes.Use(middlewares2.LogRequest)

}
