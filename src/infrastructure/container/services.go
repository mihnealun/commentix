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
	GetUserService() service.User
	GetTargetService() service.Target
	GetAppService() service.App
}

type container struct {
	config    *Config
	ogmConfig gogm.Config
	gogm      *gogm.Gogm
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
		err = instance.InitStorageDriver()
		if err != nil {
			panic(err)
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

func (c *container) InitStorageDriver() error {
	var err error

	c.gogm, err = gogm.New(&c.ogmConfig, gogm.UUIDPrimaryKeyStrategy, &entity.Comment{}, &entity.Target{}, &entity.User{}, &entity.App{})
	if err != nil {
		panic(err)
	}

	gogm.SetGlobalGogm(c.gogm)

	return nil
}

func (c *container) GetCommentService() service.Comment {
	return storage.NewCommentService(c.gogm)
}

func (c *container) GetUserService() service.User {
	return storage.NewUserService(c.gogm)
}

func (c *container) GetTargetService() service.Target {
	return storage.NewTargetService(c.gogm)
}

func (c *container) GetAppService() service.App {
	return storage.NewAppService(c.gogm)
}
