package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"todoApp-backend/src/Core/infrastructure/Web/responses"
)

func (h *handler) HomePage(c echo.Context) error {

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := struct {
		AppName string
	}{
		AppName: appName,
	}

	return c.Render(200, "home.html", data)
}

func (h *handler) LoginPage(c echo.Context) error {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.Render(200, "login.html", appName)
}

func (h *handler) SignUpPage(c echo.Context) error {

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := struct {
		AppName string
	}{
		AppName: appName,
	}

	return c.Render(200, "signup.html", data)
}

func (h *handler) FaqPage(c echo.Context) error {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.Render(200, "", appName)
}

func (h *handler) DashboardPage(c echo.Context) error {

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.Render(200, "", appName)
}
