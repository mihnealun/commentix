package service

import (
	"github.com/mihnealun/commentix/domain/entity"
)

type App interface {
	Add(app entity.App) *entity.App
	Update(app entity.App) *entity.App
	Delete(app entity.App) bool
	List() []*entity.App
}
