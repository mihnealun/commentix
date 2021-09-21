package controller

import (
	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/infrastructure/container"
	"github.com/mihnealun/commentix/infrastructure/response"
	"github.com/mindstand/gogm/v2"
	"net/http"
)

type Comment struct{}

// Get will process the input parameters and return a CommentResponse
func (pc Comment) Get(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewCommentResponse(&entity.Comment{}))
}

// List will process the input parameters and return a CommentResponse
func (pc Comment) List(context echo.Context, c container.Container) error {
	comments := c.GetCommentService().List(context.Param("target"))

	return context.JSON(http.StatusOK, response.NewCommentListResponse(comments))
}

// Create will process the input parameters and return a Comment
func (pc Comment) Create(context echo.Context, c container.Container) error {
	comment := entity.Comment{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Body:         context.FormValue("body"),
		Type:         "comment",
		Status:       "active",
	}

	result := c.GetCommentService().Create(
		context.FormValue("user"),
		context.FormValue("target"),
		context.FormValue("app"),
		comment,
	)

	if result == nil {
		return context.JSON(http.StatusNotFound, response.NewErrorResponse("Entity not found."))
	}

	return context.JSON(http.StatusCreated, response.NewCommentResponse(result))
}
