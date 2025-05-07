package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"todoApp-backend/src/internal/infrastructure/Web/responses"
)

func (h *handler) HomePage(c echo.Context) error {

	PageTitle := os.Getenv("APP_PAGE_TITLE")
	if PageTitle == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.Render(200, "", PageTitle)
}

func (h *handler) LoginPage(c echo.Context) error {
	PageTitle := os.Getenv("APP_PAGE_TITLE")
	if PageTitle == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.Render(200, "", PageTitle)
}

func (h *handler) RegisterPage(c echo.Context) error {

	PageTitle := os.Getenv("APP_PAGE_TITLE")
	if PageTitle == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.Render(200, "", PageTitle)
}

func (h *handler) DashboardPage(c echo.Context) error {

	PageTitle := os.Getenv("APP_PAGE_TITLE")
	if PageTitle == "" {
		response := responses.NewResponse("error", "error while serving Page", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.Render(200, "", PageTitle)
}
