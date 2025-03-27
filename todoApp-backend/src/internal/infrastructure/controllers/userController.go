package controllers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todoApp-backend/src/internal/infrastructure/controllers/DTO"
)

func (h *handler) createUser(c echo.Context) error {
	if c.Request().Body == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "body is required")
	}

	form := DTO.UserDTO{}

	err := c.Bind(&form)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Invalid for Data"))
	}

	err = h.UserServices.CreateUser(&form)
	if err != nil {
		errStr := fmt.Sprintf("Error while creating user: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, errors.New(errStr))
	}

	return c.JSON(http.StatusCreated, form)
}
