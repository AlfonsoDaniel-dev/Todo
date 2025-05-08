package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
)

func LogRequest(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Request received")
		return handler(c)
	}
}
