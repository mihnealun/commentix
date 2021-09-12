package container

import "C"
import (
	"context"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/domain/service"
	"github.com/mihnealun/commentix/infrastructure/storage"
	"github.com/mindstand/gogm/v2"
	"sync"
)

// Container interface that described what services it holds
type Container interface {
	GetConfig() *Config
	GetLogger(ctx context.Context) (Logger, error)
	GetCommentService() service.Comment
}

type container struct {
	config    *Config
	ogmConfig gogm.Config
}

var instance *container
var once sync.Once

// GetInstance return the container as a singleton instance
func GetInstance() (c Container, err error) {
	once.Do(func() {
		instance = &container{}
		instance.config, err = getConfigInstance()
		if err != nil {
			return
		}
		instance.ogmConfig = gogm.Config{
			Host:     instance.config.NeoHost,
			Port:     instance.config.NeoPort,
			Username: instance.config.NeoUser,
			LogLevel: instance.config.NeoLogLevel,
			Password: instance.config.NeoPass,
			PoolSize: instance.config.NeoPoolSize,
			// Encrypted:     false,
			IndexStrategy: gogm.IGNORE_INDEX,
		}
	})

	return instance, err
}

// GetConfig is returning the Config instance
func (c *container) GetConfig() *Config {
	return c.config
}

func (c *container) GetLogger(ctx context.Context) (Logger, error) {
	return newStdLogger(ctx, c.config)
}

func (c *container) GetCommentService() service.Comment {
	_gogm, err := gogm.New(&c.ogmConfig, gogm.UUIDPrimaryKeyStrategy, &entity.Comment{}, &entity.Target{}, &entity.User{}, &entity.App{})
	if err != nil {
		panic(err)
	}

	gogm.SetGlobalGogm(_gogm)

	return storage.NewCommentService(_gogm)
}
