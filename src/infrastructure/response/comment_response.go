package response

import (
	"github.com/mihnealun/commentix/domain/entity"
)

type CommentResponse struct {
	ID       string `json:"id"`
	Body     string `json:"body"`
	UserName string `json:"username"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Target   struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"target"`
	User struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"user"`
	AppName    string            `json:"app"`
	Likes      int               `json:"likes"`
	Dislikes   int               `json:"dislikes"`
	Reports    int               `json:"reports"`
	ReplyCount int               `json:"reply_count"`
	Replies    []CommentResponse `json:"replies"`
	CreatedAt  int64             `json:"created_at"`
	UpdatedAt  int64             `json:"updated_at"`
}

func NewCommentResponse(comment *entity.Comment) CommentResponse {
	if comment == nil {
		return CommentResponse{}
	}

	result := CommentResponse{
		ID:       comment.UUID,
		Body:     comment.Body,
		UserName: comment.User.Name,
		Status:   comment.Status,
		Type:     comment.Type,
		User: struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Name   string `json:"name"`
			Status string `json:"status"`
		}{
			ID:     comment.User.UUID,
			Type:   comment.User.Type,
			Name:   comment.User.Name,
			Status: comment.User.Status,
		},
		AppName:    comment.App.Name,
		Replies:    NewCommentListResponse(comment.Replies).Comments,
		Likes:      len(comment.Likers),
		Dislikes:   len(comment.Dislikers),
		Reports:    len(comment.Reporters),
		ReplyCount: len(comment.Replies),
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
	}

	if comment.Target != nil {
		result.Target = struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		}{
			ID:   comment.Target.UUID,
			Type: comment.Target.Type,
		}
	}

	return result
}
