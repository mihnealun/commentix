package service

import (
	"github.com/mihnealun/commentix/domain/entity"
)

type Comment interface {
	List(TargetID string) []*entity.Comment
	Like(CommentID string, UserID string) bool
	Dislike(CommentID string, UserID string) bool
	Report(CommentID string, UserID string) bool

	Create(UserId, TargetId, AppId string, comment entity.Comment) *entity.Comment
	Delete(CommentId string) error
	Update(CommentId string, comment entity.Comment) error
	AddReply(UserId, ParentId string, comment entity.Comment) *entity.Comment
}
