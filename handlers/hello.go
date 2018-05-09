package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handler
func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
