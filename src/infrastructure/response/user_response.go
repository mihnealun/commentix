package response

import "github.com/mihnealun/commentix/domain/entity"

type UserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func NewUserResponse(user *entity.User) UserResponse {
	return UserResponse{
		ID:     user.UUID,
		Name:   user.Name,
		Status: user.Status,
	}
}
