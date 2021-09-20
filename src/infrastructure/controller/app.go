package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/infrastructure/container"
	"github.com/mihnealun/commentix/infrastructure/response"
	"github.com/mindstand/gogm/v2"
	"net/http"
)

type App struct{}

func (a App) Get(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewAppResponse(&entity.App{}))
}

func (a App) Create(context echo.Context, c container.Container) error {
	app := entity.App{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         context.FormValue("name"),
		Slug:         context.FormValue("slug"),
	}

	result := c.GetAppService().Add(app)

	return context.JSON(http.StatusOK, response.NewAppResponse(result))
}

func (a App) Update(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewAppResponse(&entity.App{}))
}

func (a App) Delete(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewSuccessResponse("App deleted."))
}

func (a App) List(context echo.Context, c container.Container) error {
	apps := c.GetAppService().List()

	return context.JSON(http.StatusOK, response.NewAppListResponse(apps))
}
