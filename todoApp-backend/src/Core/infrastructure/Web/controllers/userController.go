package controllers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todoApp-backend/src/Core/domain"
	"todoApp-backend/src/Core/infrastructure/Web/DTO"
	"todoApp-backend/src/Core/infrastructure/Web/responses"
	"todoApp-backend/src/Core/infrastructure/auth"
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

	token := c.Request().Header.Get("authorization")
	if token == "" {
		response := responses.NewResponse("error", "couldn't read user token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	email, err := auth.GetEmailFromToken(token)
	if err != nil {
		response := responses.NewResponse("error", "couldn't read user email", nil)
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

func (h *handler) UpdateName(c echo.Context) error {

	token := c.Request().Header.Get("authorization")

	email, err := auth.GetEmailFromToken(token)
	if err != nil {
		response := responses.NewResponse("error", "couldn't read user email", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	form := DTO.UpdateUserName{}

	err = c.Bind(&form)
	if err != nil {
		response := responses.NewResponse("error", "invalid body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.UserServices.UpdateUserName(form, email)
	if errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "user not found", nil)
		return c.JSON(http.StatusUnauthorized, response)
	} else if err != nil && !errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "couldn't update user name", nil)
		return c.JSON(http.StatusOK, response)
	}

	response := responses.NewResponse("ok", "user updated successfully", nil)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateEmail(c echo.Context) error {
	token := c.Request().Header.Get("authorization")
	email, err := auth.GetEmailFromToken(token)
	if err != nil {
		response := responses.NewResponse("error", "couldn't read user email", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	form := DTO.UpdateUserEmail{}
	err = c.Bind(&form)
	if err != nil {
		response := responses.NewResponse("error", "invalid body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.UserServices.UpdateUserEmail(form, email)
	if errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "user not found", nil)
		return c.JSON(http.StatusUnauthorized, response)
	} else if err != nil && !errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "couldn't update user email", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse("ok", "user updated successfully", nil)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdatePassword(c echo.Context) error {
	token := c.Request().Header.Get("authorization")
	email, err := auth.GetEmailFromToken(token)
	if err != nil {
		response := responses.NewResponse("error", "couldn't read user email", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	form := DTO.UpdateUserPassword{}
	err = c.Bind(&form)
	if err != nil {
		response := responses.NewResponse("error", "invalid body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.UserServices.UpdateUserPassword(form, email)
	if errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "user not found", nil)
		return c.JSON(http.StatusUnauthorized, response)
	} else if err != nil && !errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "couldn't update user password", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse("ok", "user updated successfully", nil)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteUser(c echo.Context) error {
	token := c.Request().Header.Get("authorization")
	email, err := auth.GetEmailFromToken(token)
	if err != nil {
		response := responses.NewResponse("error", "couldn't read user email", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	form := DTO.DeleteUser{}
	err = c.Bind(&form)
	if err != nil {
		response := responses.NewResponse("error", "invalid body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.UserServices.DeleteUser(form, email)
	if errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "user not found", nil)
		return c.JSON(http.StatusBadRequest, response)
	} else if err != nil && !errors.Is(err, domain.ErrNotFound) {
		response := responses.NewResponse("error", "couldn't delete user data", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse("ok", "user deleted successfully", nil)
	return c.JSON(http.StatusOK, response)
}
