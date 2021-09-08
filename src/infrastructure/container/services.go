package container

import (
	"context"
	"github.com/mihnealun/commentix/domain/service"
	"sync"
)

// Container interface that described what services it holds
type Container interface {
	GetConfig() *Config
	GetLogger(ctx context.Context) (Logger, error)
	GetCommentService() service.Comment
	InitStorage()
}

type container struct {
	config *Config
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
	c.InitStorage()
	return service.NewComment()
}

func (c *container) InitStorage() {
	//return service.NewComment()
}
