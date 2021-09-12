package service

import (
	"github.com/mihnealun/commentix/domain/entity"
)

const (
	DaysInMonth = 30
	DaysInYear  = 360
)

type Comment interface {
	Like(CommentID int, UserID int) error
	Dislike(CommentID int, UserID int) error
	Report(CommentID int, UserID int) error

	Add(UserId, TargetId, AppId string, comment entity.Comment) *entity.Comment
	Delete(CommentId string) error
	Update(CommentId string, comment entity.Comment) error
	Reply(UserId, ParentId string, comment entity.Comment) string
}
