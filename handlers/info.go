package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/tgracchus/echo-app/repo"
)

func NewInfoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		version, err := repo.Info(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, fmt.Sprintf("Connected to %s:", version))
	}
}
