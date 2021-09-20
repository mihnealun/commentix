package storage

import (
	"context"
	"github.com/mihnealun/commentix/domain/entity"
	"github.com/mihnealun/commentix/domain/service"
	"github.com/mindstand/gogm/v2"
	"log"
)

type comment struct {
	driver *gogm.Gogm
}

func NewCommentService(driver *gogm.Gogm) service.Comment {
	return &comment{
		driver: driver,
	}
}

func (c *comment) AddComment(UserId, TargetId, AppId string, comment entity.Comment) *entity.Comment {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer c.commitAndClose(sess)

	comment.Target = c.GetTarget(TargetId)
	comment.App = c.GetApp(AppId)
	comment.User = c.GetUser(UserId)

	err = sess.SaveDepth(context.Background(), &comment, 2)
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

func (c *comment) AddCommentRaw(User *entity.User, Target *entity.Target, App *entity.App, Comment entity.Comment) *entity.Comment {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer c.commitAndClose(sess)

	Comment.Target = Target
	Comment.App = App
	Comment.User = User

	err = sess.SaveDepth(context.Background(), &Comment, 2)
	if err != nil {
		panic(err)
	}

	var result entity.Comment
	err = sess.Load(context.Background(), &result, Comment.UUID)
	if err != nil {
		panic(err)
	}

	return &result
}

func (c *comment) AddApp(app entity.App) *entity.App {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer c.commitAndClose(sess)

	err = sess.SaveDepth(context.Background(), app, 2)
	if err != nil {
		panic(err)
	}

	var result entity.App
	err = sess.Load(context.Background(), &result, app.UUID)
	if err != nil {
		panic(err)
	}

	return &result
}

func (c *comment) AddTarget(target entity.Target) *entity.Target {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer c.commitAndClose(sess)

	err = sess.SaveDepth(context.Background(), target, 2)
	if err != nil {
		panic(err)
	}

	var result entity.Target
	err = sess.Load(context.Background(), &result, target.UUID)
	if err != nil {
		panic(err)
	}

	return &result
}

func (c *comment) DeleteComment(CommentId string) error {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		return err
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer c.commitAndClose(sess)

	var comment entity.Comment
	err = sess.Load(context.Background(), &comment, &CommentId)
	if err != nil {
		return err
	}

	comment.Status = "deleted"

	err = sess.SaveDepth(context.Background(), comment, 1)
	if err != nil {
		return err
	}

	return nil
}

func (c *comment) UpdateComment(CommentId string, comment entity.Comment) error {
	panic("implement me")
}

func (c *comment) ListComments(TargetID string) []*entity.Comment {
	/*
		MATCH (u:User)-[r1:CREATED]-(n:Comment)-[r:TARGETS]-(t:Target {uuid:'c62f74df-a965-4468-b0ed-d18919db2403'}) RETURN u, r1, n, r, t LIMIT 25
	*/

	var allComments []*entity.Comment

	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Println(err.Error())
		return allComments
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Println(err.Error())
		return allComments
	}

	defer c.commitAndClose(sess)

	err = sess.LoadAll(context.Background(), &allComments)
	if err != nil {
		log.Println(err.Error())
	}

	return allComments
}

func (c *comment) Like(CommentID string, UserID string) bool {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer c.commitAndClose(sess)

	comment := c.GetComment(CommentID)
	user := c.GetUser(UserID)

	if comment.Likers != nil {
		for _, liker := range comment.Likers {
			if liker.UUID == user.UUID {
				// Already liked
				return false
			}
		}
	}

	comment.Likers = append(comment.Likers, user)

	err = sess.SaveDepth(context.Background(), user, 2)
	if err != nil {
		panic(err)
	}

	return true
}

func (c *comment) Dislike(CommentID string, UserID string) bool {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer c.commitAndClose(sess)

	comment := c.GetComment(CommentID)
	user := c.GetUser(UserID)

	if comment.Dislikers != nil {
		for _, disliker := range comment.Dislikers {
			if disliker.UUID == user.UUID {
				// Already disliked
				return false
			}
		}
	}

	comment.Dislikers = append(comment.Dislikers, user)

	err = sess.SaveDepth(context.Background(), user, 2)
	if err != nil {
		panic(err)
	}

	return true
}

func (c *comment) Report(CommentID string, UserID string) bool {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer c.commitAndClose(sess)

	comment := c.GetComment(CommentID)
	user := c.GetUser(UserID)

	if comment.Reporters != nil {
		for _, reporter := range comment.Reporters {
			if reporter.UUID == user.UUID {
				// Already reported
				return false
			}
		}
	}

	comment.Reporters = append(comment.Reporters, user)

	err = sess.SaveDepth(context.Background(), user, 2)
	if err != nil {
		panic(err)
	}

	return true
}

func (c *comment) AddReply(UserId, ParentId string, comment entity.Comment) *entity.Comment {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer c.commitAndClose(sess)

	comment.Type = "reply"
	parent := c.GetComment(ParentId)
	user := c.GetUser(UserId)
	reply := c.AddCommentRaw(user, parent.Target, parent.App, comment)
	parent.Replies = append(parent.Replies, reply)

	err = sess.SaveDepth(context.Background(), parent, 2)
	if err != nil {
		panic(err)
	}

	return reply
}

func (c *comment) GetUser(UserId string) *entity.User {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}

	defer sess.Close()

	user := &entity.User{}

	err = sess.Load(context.Background(), user, UserId)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return user
}

func (c *comment) GetComment(CommentId string) *entity.Comment {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}

	defer sess.Close()

	comment := &entity.Comment{}

	err = sess.Load(context.Background(), comment, CommentId)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return comment
}

func (c *comment) GetTarget(TargetId string) *entity.Target {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}

	defer sess.Close()

	target := &entity.Target{}
	err = sess.Load(context.Background(), target, TargetId)

	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return target
}

func (c *comment) GetApp(AppId string) *entity.App {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer c.commitAndClose(sess)

	app := &entity.App{}
	//query := `MATCH p=(ap:App {uuid:$appid}) RETURN p`
	err = sess.Load(context.Background(), app, &AppId)

	//err = sess.Query(context.Background(), query, map[string]interface{}{"appid": AppId}, app)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return app
}

func (c *comment) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
