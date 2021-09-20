package service

import (
	"github.com/mihnealun/commentix/domain/entity"
)

const (
	DaysInMonth = 30
	DaysInYear  = 360
)

type Comment interface {
	ListComments(TargetID string) []*entity.Comment
	Like(CommentID string, UserID string) bool
	Dislike(CommentID string, UserID string) bool
	Report(CommentID string, UserID string) bool

	AddComment(UserId, TargetId, AppId string, comment entity.Comment) *entity.Comment
	DeleteComment(CommentId string) error
	UpdateComment(CommentId string, comment entity.Comment) error
	AddReply(UserId, ParentId string, comment entity.Comment) *entity.Comment

	AddTarget(target entity.Target) *entity.Target
	AddApp(app entity.App) *entity.App
}
