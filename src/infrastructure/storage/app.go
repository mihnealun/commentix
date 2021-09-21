package storage

import (
	"context"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/domain/service"
	"github.com/mindstand/gogm/v2"
	"log"
)

type app struct {
	driver *gogm.Gogm
}

func NewAppService(driver *gogm.Gogm) service.App {
	return &app{
		driver: driver,
	}
}

func (a *app) Add(app entity.App) *entity.App {
	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Fatal(err)
	}
	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer a.commitAndClose(sess)

	err = sess.SaveDepth(context.Background(), &app, 1)
	if err != nil {
		log.Fatal(err)
	}

	var result entity.App
	err = sess.Load(context.Background(), &result, app.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (a *app) Update(app entity.App) *entity.App {
	return &app
}

func (a *app) Delete(app entity.App) bool {
	return true
}

func (a *app) List() []*entity.App {
	var allApps []*entity.App

	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Println(err.Error())
		return allApps
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Println(err.Error())
		return allApps
	}

	defer a.commitAndClose(sess)

	err = sess.LoadAll(context.Background(), &allApps)
	if err != nil {
		log.Println(err.Error())
	}

	return allApps
}

func (a *app) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
