package storage

import (
	"context"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/domain/service"
	"github.com/mindstand/gogm/v2"
	"log"
)

type user struct {
	driver *gogm.Gogm
}

func NewUserService(driver *gogm.Gogm) service.User {
	return &user{
		driver: driver,
	}
}

func (u *user) Add(user entity.User) *entity.User {
	sess, err := u.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Fatal(err)
	}

	defer u.commitAndClose(sess)

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = sess.SaveDepth(context.Background(), &user, 1)
	if err != nil {
		log.Fatal(err)
	}

	var result entity.User
	err = sess.Load(context.Background(), &result, user.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (u *user) Update(user entity.User) *entity.User {
	return &user
}

func (u *user) Delete(user entity.User) bool {
	return true
}

func (u *user) List() []*entity.User {
	var allUsers []*entity.User

	sess, err := u.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		log.Println(err.Error())
		return allUsers
	}

	defer sess.Close()

	err = sess.LoadAll(context.Background(), &allUsers)
	if err != nil {
		log.Println(err.Error())
	}

	return allUsers
}

func (u *user) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
