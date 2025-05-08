package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todoApp-backend/src/Core/infrastructure/Web/responses"
	"todoApp-backend/src/Core/infrastructure/auth"
)

func AuthMiddleWare(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")
		if token == "" {
			response := responses.NewResponse("error", "authorization header is required", nil)
			c.JSON(http.StatusBadRequest, response)
		}
		_, err := auth.ValidateToken(token)
		if err != nil {
			if err = c.Redirect(http.StatusUnauthorized, "/login"); err != nil {
				response := responses.NewResponse("error", "unauthorized", nil)
				return c.JSON(http.StatusUnauthorized, response)
			}
		}

		return f(c)
	}
}
