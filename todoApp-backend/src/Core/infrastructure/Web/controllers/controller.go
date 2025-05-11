package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"path/filepath"
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

	return &handler{
		UserServices: *userService,
		TaskServices: *taskService,
	}
}

type controller struct {
	E              *echo.Echo
	handlers       *handler
	Renderer       echo.Renderer
	AppRoutes      *echo.Group
	ApiGroups      *echo.Group
	UserRepository domain.UserRepository
	TaskRepository domain.TaskRepository
}

func NewController(db *sql.DB, echo *echo.Echo, templatesDir string) {

	userRepository := data.NewUserRepository(db)

	taskRepository := data.NewTaskRepository(db)

	handlerInstance := newHandler(userRepository, taskRepository)

	apiGroups := echo.Group("/api/v1")

	templateEngine := newTemplates(templatesDir)

	controllerInstance := &controller{
		E:              echo,
		handlers:       handlerInstance,
		ApiGroups:      apiGroups,
		Renderer:       templateEngine,
		UserRepository: userRepository,
		TaskRepository: taskRepository,
	}

	controllerInstance.E.Renderer = controllerInstance.Renderer
	controllerInstance.E.Static("/static", filepath.Join(templatesDir, "static"))

	controllerInstance.MountEndpoints()

}

func (C *controller) MountEndpoints() {

	C.UserRoutes()
	C.AppPublicRoutes()
	C.TaskRoutes()
}

func (C *controller) AppPublicRoutes() {

	appPublicRoutes := C.E.Group("")

	appPublicRoutes.Use(middlewares2.LogRequest)
	appPublicRoutes.GET("/home", C.handlers.HomePage)
	appPublicRoutes.GET("/login", C.handlers.LoginPage)
	appPublicRoutes.GET("/signup", C.handlers.SignUpPage)
	appPublicRoutes.GET("/faq", C.handlers.FaqPage)
	appPublicRoutes.GET("pricing", C.handlers.PricingPage)

	C.AppRoutes = appPublicRoutes
}

func (C *controller) UserRoutes() {

	PublicRoutes := C.ApiGroups.Group("/user")

	PublicRoutes.Use(middlewares2.LogRequest)

	PublicRoutes.GET("/create", C.handlers.CreateUser)
	PublicRoutes.POST("/login", C.handlers.Login)
	PublicRoutes.POST("/login/google/:code", C.handlers.LoginOauth)

	userPrivateRoutes := C.ApiGroups.Group("/user/private")

	userPrivateRoutes.Use(middlewares2.AuthMiddleWare)

}

func (C *controller) TaskRoutes() {
	taskPublicRoutes := C.ApiGroups.Group("/task")
	taskPublicRoutes.Use(middlewares2.LogRequest)

	taskPublicRoutes.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

}
