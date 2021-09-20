package response

import "github.com/mihnealun/commentix/domain/entity"

type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

func NewUserListResponse(users []*entity.User) UserListResponse {
	result := UserListResponse{
		Users: []UserResponse{},
	}

	for _, u := range users {
		result.Users = append(result.Users, NewUserResponse(u))
	}

	return result
}
