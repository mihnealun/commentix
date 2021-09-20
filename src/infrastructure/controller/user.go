package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/infrastructure/container"
	"github.com/mihnealun/commentix/infrastructure/response"
	"github.com/mindstand/gogm/v2"
	"net/http"
)

type User struct{}

func (pc User) Get(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewUserResponse(&entity.User{}))
}

func (pc User) Create(context echo.Context, c container.Container) error {
	user := entity.User{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         context.FormValue("name"),
		Type:         "user",
		Status:       "active",
	}

	result := c.GetUserService().Add(user)

	return context.JSON(http.StatusOK, response.NewUserResponse(result))
}

func (pc User) Update(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewUserResponse(&entity.User{}))
}

func (pc User) Delete(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewSuccessResponse("User deleted."))
}

func (pc User) List(context echo.Context, c container.Container) error {
	users := c.GetUserService().List()

	return context.JSON(http.StatusOK, response.NewUserListResponse(users))
}
