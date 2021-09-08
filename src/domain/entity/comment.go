package entity

type User struct {
	ID         int    `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	PlatformID string `json:"platform_id" bson:"platform_id"`
}

type Status struct {
	ID    int    `json:"id" bson:"id"`
	Label string `json:"label" bson:"label"`
}

type Target struct {
	ID   int    `json:"id" bson:"id"`
	Type string `json:"type" bson:"type"`
}

type App struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Slug string `json:"slug" bson:"slug"`
}

// Comment document structure
type Comment struct {
	CommentID int       `json:"comment_id" bson:"comment_id"`
	Body      string    `json:"body" bson:"body"`
	User      User      `json:"user" bson:"user"`
	Status    Status    `json:"status" bson:"status"`
	Target    Target    `json:"target" bson:"target"`
	App       App       `json:"app" bson:"app"`
	Replies   []Comment `json:"replies" bson:"replies"`
}
