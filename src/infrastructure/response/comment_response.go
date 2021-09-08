package response

import "github.com/mihnealun/commentix/domain/entity"

type CommentResponse struct {
	ID       string `json:"id"`
	Body     string `json:"body"`
	UserName string `json:"user"`
	Status   string `json:"status"`
	Target   struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"target"`
	AppName string `json:"app"`
}

func NewCommentResponse(comment *entity.Comment) CommentResponse {
	return CommentResponse{
		Body:     comment.Body,
		UserName: comment.User.Name,
		Status:   comment.Status.Label,
		Target: struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
		}{
			ID:   comment.Target.ID,
			Type: comment.Target.Type,
		},
		AppName: comment.App.Name,
	}
}
