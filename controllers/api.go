package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Get(c echo.Context) error {
	return c.String(http.StatusOK, "It's Alive!")
}
