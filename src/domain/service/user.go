package service

import (
	"github.com/mihnealun/commentix/domain/entity"
)

type User interface {
	Add(user entity.User) *entity.User
	Update(user entity.User) *entity.User
	Delete(user entity.User) bool
	List() []*entity.User
}
