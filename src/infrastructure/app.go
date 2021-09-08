package infrastructure

import (
	"fmt"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/mihnealun/commentix/infrastructure/container"
	"github.com/mihnealun/commentix/infrastructure/http/echo/handler"
	"github.com/mihnealun/commentix/infrastructure/http/echo/middleware"
	"github.com/mihnealun/commentix/infrastructure/route"
)

// Start method is bootstrapping and starting the entire application
func Start(containerInstance container.Container) error {
	e := echo.New()
	config := containerInstance.GetConfig()

	e.HTTPErrorHandler = handler.HTTPErrorHandler

	e.Use(echoMiddleware.Recover())
	e.Use(middleware.Logger(containerInstance))
	e.Use(echoMiddleware.Gzip())

	route.PreparePublicRoutes(e, containerInstance)

	address := fmt.Sprintf("%s:%d", config.Interface, config.Port)
	err := e.Start(address)
	if err != nil {
		return err
	}

	return nil
}
