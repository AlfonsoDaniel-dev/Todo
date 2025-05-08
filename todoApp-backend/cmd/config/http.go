package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"todoApp-backend/src/Core/infrastructure/Web/controllers"
)

func NewHttp(templatesDir string) *echo.Echo {

	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"192.168.5.95"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"authorization", "content-type"},
		AllowCredentials: true,
	}))

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())

	templateEngine := newTemplates(templatesDir)
	server.Renderer = templateEngine

	controller := controllers.NewController()

	return server
}
