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

const (
	DateFormat = "2006-01-02T15:04:05Z"
)

type Comment struct{}

// Get will process the input parameters and return a CommentResponse
func (pc Comment) Get(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewCommentResponse(&entity.Comment{}))
}

// List will process the input parameters and return a CommentResponse
func (pc Comment) List(context echo.Context, c container.Container) error {
	comments := c.GetCommentService().ListComments(context.Param("target"))

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

	result := c.GetCommentService().AddComment(
		context.FormValue("user"),
		context.FormValue("target"),
		context.FormValue("app"),
		comment)

	return context.JSON(http.StatusOK, response.NewCommentResponse(result))
}
