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

func (c *comment) Create(UserId, TargetId, AppId string, comment entity.Comment) *entity.Comment {
	sess, err := c.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Fatal(err)
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer c.commitAndClose(sess)

	comment.Target = c.GetTarget(TargetId)
	comment.App = c.GetApp(AppId)
	comment.User = c.GetUser(UserId)

	if comment.Target == nil || comment.App == nil || comment.User == nil {
		log.Println("Entity not found")
		return nil
	}

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

func (c *comment) AddRaw(User *entity.User, Target *entity.Target, App *entity.App, Comment entity.Comment) *entity.Comment {
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

func (c *comment) Delete(CommentId string) error {
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

func (c *comment) List(TargetID string) (comments []*entity.Comment) {
	t := c.GetTarget(TargetID)

	if t == nil || len(t.Comments) == 0 {
		log.Println("No comments found")
		return comments
	}

	return t.Comments
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
	//if comment.User.UUID == UserID {
	//	return false
	//}
	user := c.GetUser(UserID)

	if comment.Likers != nil {
		for _, liker := range comment.Likers {
			if liker.UUID == user.UUID {
				// Already liked
				return false
			}
		}
	}

	user.LikedComments = append(user.LikedComments, comment)
	comment.Likers = append(comment.Likers, user)

	err = sess.SaveDepth(context.Background(), comment, 2)
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
	//if comment.User.UUID == UserID {
	//	return false
	//}

	user := c.GetUser(UserID)

	if comment.Dislikers != nil {
		for _, disliker := range comment.Dislikers {
			if disliker.UUID == user.UUID {
				// Already disliked
				return false
			}
		}
	}

	user.DislikedComments = append(user.DislikedComments, comment)
	comment.Dislikers = append(comment.Dislikers, user)

	err = sess.SaveDepth(context.Background(), comment, 2)
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
	//if comment.User.UUID == UserID {
	//	return false
	//}
	user := c.GetUser(UserID)

	if comment.Reporters != nil {
		for _, reporter := range comment.Reporters {
			if reporter.UUID == user.UUID {
				// Already reported
				return false
			}
		}
	}

	user.ReportedComments = append(user.ReportedComments, comment)
	comment.Reporters = append(comment.Reporters, user)

	err = sess.SaveDepth(context.Background(), comment, 2)
	if err != nil {
		panic(err)
	}

	return true
}

func (c *comment) Reply(UserId, ParentId string, comment entity.Comment) *entity.Comment {
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
	reply := c.AddRaw(user, nil, parent.App, comment)
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

	err = sess.LoadDepth(context.Background(), target, TargetId, 3)
	if err != nil {
		log.Println(err.Error())
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

	err = sess.Load(context.Background(), app, &AppId)
	if err != nil {
		log.Println(err.Error())
		return app
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
