package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var (
	echoRouter *echo.Echo
)

func (a *api) routes() *echo.Echo {

	// ensure we are only creating the router once, regardless of how many times .routes() is invoked
	if echoRouter != nil {
		return echoRouter
	}

	echoRouter := echo.New()
	echoRouter.Logger.SetLevel(log.OFF)

	// Middleware
	echoRouter.Use(middleware.Logger())
	echoRouter.Use(middleware.Recover())

	// Route => handler
	echoRouter.GET("/", func(c echo.Context) error {
		a.Log.Infof("hello world")
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	return echoRouter
}
