package controllers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/controllers/DTO"
	"todoApp-backend/src/internal/infrastructure/responses"
)

func (h *handler) CreateUser(c echo.Context) error {
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

func (h *handler) GetUser(c echo.Context) error {

	email := c.Request().Header.Get("authorization")
	if email == "" {
		response := responses.NewResponse("error", "couldn't read user token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	user, err := h.UserServices.GetUser(email)
	if errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "user not found", nil)
		return c.JSON(http.StatusUnauthorized, response)
	} else if err != nil && !errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "couldn't get user data", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse("ok", "user obtained successfully", user)
	return c.JSON(http.StatusOK, response)
}
