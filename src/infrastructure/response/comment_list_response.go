package response

import (
	"github.com/mihnealun/commentix/domain/entity"
)

type CommentListResponse struct {
	Comments []CommentResponse `json:"comments"`
}

func NewCommentListResponse(comments []*entity.Comment) CommentListResponse {
	result := CommentListResponse{
		Comments: []CommentResponse{},
	}

	if comments == nil {
		return result
	}

	for _, c := range comments {
		result.Comments = append(result.Comments, NewCommentResponse(c))
	}

	return result
}
