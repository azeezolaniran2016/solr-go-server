package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	echoRouter.GET("/healthcheck", a.healthcheck)

	echoRouter.GET("/query", a.query)

	return echoRouter
}
