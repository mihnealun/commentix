package route

import (
	"github.com/labstack/echo/v4"

	"github.com/mihnealun/commentix/infrastructure/container"
	"github.com/mihnealun/commentix/infrastructure/controller"
)

// PreparePublicRoutes add the necessary public routes to echo
func PreparePublicRoutes(e *echo.Echo, containerInstance container.Container) {
	e.GET("/comment", func(c echo.Context) error {
		commentController := controller.Comment{}

		return commentController.List(c, containerInstance)
	})

	e.GET("/comment/:id", func(c echo.Context) error {
		commentController := controller.Comment{}

		return commentController.Get(c, containerInstance)
	})

	e.POST("/comment", func(c echo.Context) error {
		ctrl := controller.Comment{}

		return ctrl.Create(c, containerInstance)
	})
}
