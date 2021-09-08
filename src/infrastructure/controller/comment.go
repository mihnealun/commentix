package controller

import (
	"net/http"
	"strconv"

	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/infrastructure/container"
	"github.com/mihnealun/commentix/infrastructure/response"
)

const (
	DateFormat = "2006-01-02T15:04:05Z"
)

type Comment struct{}

// Get will process the input parameters and return a CommentResponse
func (pc Comment) Get(context echo.Context, c container.Container) error {
	//commentID, err := strconv.Atoi(context.Param("id"))
	//if err != nil {
	//	log.Error(err.Error())
	//	return echo.ErrBadRequest.SetInternal(err)
	//}
	//
	// 	commentService := c.GetCommentService()
	// 	comment, err := commentService.Get(commentID, c.GetCommentsCollection())
	// 	if err != nil {
	// 		log.Error(err.Error())
	// 		return echo.ErrNotFound.SetInternal(err)
	// 	}

	return context.JSON(http.StatusOK, response.NewCommentResponse(&entity.Comment{}))
}

// List will process the input parameters and return a CommentResponse
func (pc Comment) List(context echo.Context, c container.Container) error {

	var comments []entity.Comment

	return context.JSON(http.StatusOK, response.NewCommentListResponse(comments))
}

// Create will process the input parameters and return a Comment
func (pc Comment) Create(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewCommentResponse(&entity.Comment{}))
}

func (pc Comment) buildComment(context echo.Context) entity.Comment {
	result := entity.Comment{}

	result.CommentID, _ = strconv.Atoi(context.FormValue("id"))
	result.Body = context.FormValue("body")
	result.Status = entity.Status{
		ID:    12,
		Label: "Active",
	}
	result.User = entity.User{
		ID:         11,
		Name:       context.FormValue("user"),
		PlatformID: "comments",
	}
	result.App = entity.App{
		ID:   12,
		Name: "RealityKings",
		Slug: "rk",
	}

	return result
}
