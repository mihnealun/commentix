package service

import (
	"github.com/mihnealun/commentix/domain/entity"
)

type Target interface {
	Add(target entity.Target) *entity.Target
	Update(user entity.Target) *entity.Target
	Delete(user entity.Target) bool
	List() []*entity.Target
	Get(id string) *entity.Target
}
