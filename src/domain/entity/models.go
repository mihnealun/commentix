package entity

import (
	"github.com/mindstand/gogm/v2"
)

type User struct {
	gogm.BaseUUIDNode
	Name             string     `gogm:"name=name"`
	Type             string     `gogm:"name=type"`
	Status           string     `gogm:"name=status"`
	Comments         []*Comment `gogm:"direction=outgoing;relationship=CREATED"`
	DislikedComments []*Comment `gogm:"direction=outgoing;relationship=DISLIKED"`
	LikedComments    []*Comment `gogm:"direction=outgoing;relationship=LIKED"`
	ReportedComments []*Comment `gogm:"direction=outgoing;relationship=REPORTED"`
	DislikedTarget   []*Comment `gogm:"direction=outgoing;relationship=DISLIKED"`
	LikedTarget      []*Comment `gogm:"direction=outgoing;relationship=LIKED"`
	ReportedTarget   []*Comment `gogm:"direction=outgoing;relationship=REPORTED"`
}

type Target struct {
	gogm.BaseUUIDNode
	Name      string     `gogm:"name=name"`
	Type      string     `gogm:"name=type"`
	Url       string     `gogm:"name=url"`
	Comments  []*Comment `gogm:"direction=incoming;relationship=TARGETS"`
	Likers    []*User    `gogm:"direction=incoming;relationship=LIKED"`
	Dislikers []*User    `gogm:"direction=incoming;relationship=DISLIKED"`
	Reporters []*User    `gogm:"direction=incoming;relationship=REPORTED"`
}

type App struct {
	gogm.BaseUUIDNode
	Name     string     `gogm:"name=name"`
	Slug     string     `gogm:"name=slug"`
	Comments []*Comment `gogm:"direction=incoming;relationship=POSTED_ON"`
}

type Comment struct {
	gogm.BaseUUIDNode
	Body      string     `gogm:"name=body"`
	Type      string     `gogm:"name=type"`
	Status    string     `gogm:"name=status"`
	CreatedAt int        `gogm:"name=created_at"`
	User      *User      `gogm:"direction=incoming;relationship=CREATED"`
	Target    *Target    `gogm:"direction=outgoing;relationship=TARGETS"`
	App       *App       `gogm:"direction=outgoing;relationship=POSTED_ON"`
	Replies   []*Comment `gogm:"direction=incoming;relationship=REPLIES_TO"`
	Likers    []*User    `gogm:"direction=incoming;relationship=LIKED"`
	Dislikers []*User    `gogm:"direction=incoming;relationship=DISLIKED"`
	Reporters []*User    `gogm:"direction=incoming;relationship=REPORTED"`
	Parent    *Comment   `gogm:"direction=outgoing;relationship=REPLIES_TO"`
}
