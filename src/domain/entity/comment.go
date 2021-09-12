package entity

import (
	"github.com/mindstand/gogm/v2"
)

//type tdString string
//type tdInt int

type User struct {
	gogm.BaseUUIDNode
	Name     string     `gogm:"name=name"`
	Type     string     `gogm:"name=type"`
	Status   string     `gogm:"name=status"`
	Comments []*Comment `gogm:"direction=outgoing;relationship=commented_rel"`
}

type Target struct {
	gogm.BaseUUIDNode
	Name     string     `gogm:"name=name"`
	Type     string     `gogm:"name=type"`
	Url      string     `gogm:"name=url"`
	Comments []*Comment `gogm:"direction=incoming;relationship=targets_rel"`
}

type App struct {
	gogm.BaseUUIDNode
	Name     string     `gogm:"name=name"`
	Slug     string     `gogm:"name=slug"`
	Comments []*Comment `gogm:"direction=incoming;relationship=posted_on_rel"`
}

type Comment struct {
	gogm.BaseUUIDNode
	Body      string     `gogm:"name=body"`
	Type      string     `gogm:"name=type"`
	Status    string     `gogm:"name=status"`
	CreatedAt int        `gogm:"name=created_at"`
	User      *User      `gogm:"direction=incoming;relationship=commented_rel"`
	Target    *Target    `gogm:"direction=outgoing;relationship=targets_rel"`
	App       *App       `gogm:"direction=outgoing;relationship=posted_on_rel"`
	Replies   []*Comment `gogm:"direction=incoming;relationship=replies_to_rel"`
}
