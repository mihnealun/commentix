package storage

import (
	"context"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/domain/service"
	"github.com/mindstand/gogm/v2"
)

type comment struct {
	driver *gogm.Gogm
}

func NewCommentService(driver *gogm.Gogm) service.Comment {
	return &comment{
		driver: driver,
	}
}

func (c *comment) Add(UserId, TargetId, AppId string, comment entity.Comment) *entity.Comment {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	defer sess.Close()

	comment.Target = c.GetTarget(TargetId)
	comment.App = c.GetApp(AppId)
	comment.User = c.GetUser(UserId)

	err = sess.SaveDepth(context.Background(), comment, 2)
	if err != nil {
		panic(err)
	}

	var result entity.Comment
	err = sess.Load(context.Background(), &result, comment.UUID)
	if err != nil {
		panic(err)
	}

	return &result
}

func (c *comment) Delete(CommentId string) error {
	panic("implement me")
}

func (c *comment) Update(CommentId string, comment entity.Comment) error {
	panic("implement me")
}

func (c *comment) Like(CommentID int, UserID int) error {
	return nil
}

func (c *comment) Dislike(CommentID int, UserID int) error {
	return nil
}

func (c *comment) Report(CommentID int, UserID int) error {
	return nil
}

func (c *comment) Reply(UserId, ParentId string, comment entity.Comment) string {
	return ""
}

func (c *comment) GetUser(UserId string) *entity.User {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	user := &entity.User{}
	query := `MATCH p=(usr:User {id:$userid}) RETURN p`

	err = sess.Query(context.Background(), query, map[string]interface{}{"userid": UserId}, user)
	if err != nil {
		panic(err)
	}

	return user
}

func (c *comment) GetTarget(TargetId string) *entity.Target {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	target := &entity.Target{}
	query := `MATCH p=(trgt:Target {id:$targetid}) RETURN p`

	err = sess.Query(context.Background(), query, map[string]interface{}{"targetid": TargetId}, target)
	if err != nil {
		panic(err)
	}

	return target
}

func (c *comment) GetApp(AppId string) *entity.App {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	app := &entity.App{}
	query := `MATCH p=(ap:App {id:$appid}) RETURN p`

	err = sess.Query(context.Background(), query, map[string]interface{}{"appid": AppId}, app)
	if err != nil {
		panic(err)
	}

	return app
}
