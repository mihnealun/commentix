package storage

import (
	"context"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/domain/service"
	"github.com/mindstand/gogm/v2"
	"log"
)

type target struct {
	driver *gogm.Gogm
}

func NewTargetService(driver *gogm.Gogm) service.Target {
	return &target{
		driver: driver,
	}
}

func (t *target) Add(target entity.Target) *entity.Target {
	sess, err := t.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Fatal(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer t.commitAndClose(sess)

	err = sess.SaveDepth(context.Background(), &target, 1)
	if err != nil {
		log.Fatal(err)
	}

	var result entity.Target
	err = sess.Load(context.Background(), &result, target.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (t *target) Update(user entity.Target) *entity.Target {
	return &user
}

func (t *target) Delete(user entity.Target) bool {
	return true
}

func (t *target) List() []*entity.Target {
	var allTargets []*entity.Target

	sess, err := t.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Println(err.Error())
		return allTargets
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Println(err.Error())
		return allTargets
	}

	defer t.commitAndClose(sess)

	err = sess.LoadAll(context.Background(), &allTargets)
	if err != nil {
		log.Println(err.Error())
	}

	return allTargets
}

func (t *target) Get(id string) *entity.Target {
	target := &entity.Target{}

	sess, err := t.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		log.Println(err.Error())
		return target
	}
	defer sess.Close()

	err = sess.Load(context.Background(), target, id)

	if err != nil {
		log.Println("asdf")
		log.Println(err.Error())
	}

	return target
}

func (t *target) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
