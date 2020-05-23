package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *api) healthcheck(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, []byte(`{"status": "ok"}`))
}
