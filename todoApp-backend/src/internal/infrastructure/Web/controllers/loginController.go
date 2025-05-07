package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/Web/DTO"
	"todoApp-backend/src/internal/infrastructure/Web/responses"
	"todoApp-backend/src/internal/infrastructure/auth"
)

func (h *handler) Login(c echo.Context) error {
	form := DTO.LoginDTO{}

	err := c.Bind(&form)
	if err != nil {
		response := responses.NewResponse("error", "Invalid Data", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.UserServices.Login(form)

	if !errors.Is(err, domain.ErrNotFound) || !errors.Is(err, domain.ErrWrongPassword) {
		response := responses.NewResponse("error", "login failed", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	if errors.Is(err, domain.ErrNotFound) {

	} else if errors.Is(err, domain.ErrWrongPassword) {
		response := responses.NewResponse("error", "login failed", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	token, err := auth.GenerateToken(form.Email)
	if err != nil {
		response := responses.NewResponse("error", "error while generating token", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse("ok", "user Login Success", token)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) LoginOauth(c echo.Context) error {

	code := c.Param("code")
	if code == "" {
		response := responses.NewResponse("error", "Error while reading token", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	accesstoken, err := auth.GetAccessToken(code)
	if err != nil {
		response := responses.NewResponse("error", "Error while authorizing token", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	UserName, UserEmail, err := auth.GetUserInfo(accesstoken.AccessToken)
	if err != nil {
		response := responses.NewResponse("error", "error while getting user information", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = h.UserServices.OAuthLogin(UserName, UserEmail)
	if errors.Is(err, domain.UserAlreadyExists) {

		token, err := auth.GenerateToken(UserEmail)
		if err != nil {
			response := responses.NewResponse("error", "error while generating token", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := responses.NewResponse("ok", "user Login Success", map[string]string{"token": token})
		return c.JSON(http.StatusOK, response)
	} else if errors.Is(err, domain.ErrNotFound) {
		token, err := auth.GenerateToken(UserEmail)
		if err != nil {
			response := responses.NewResponse("error", "error while generating token", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := responses.NewResponse("ok", "user created successfully", map[string]string{"token": token})

		return c.JSON(http.StatusOK, response)
	}

	response := responses.NewResponse("error", "login failed", err)
	return c.JSON(http.StatusInternalServerError, response)
}
