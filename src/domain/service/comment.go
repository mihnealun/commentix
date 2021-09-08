package service

const (
	DaysInMonth = 30
	DaysInYear  = 360
)

// PaymentCalculator main interface used for processing payment plan processing
type Comment interface {
	Like(CommentID int, UserID int) error
	Dislike(CommentID int, UserID int) error
	Report(CommentID int, UserID int) error
}

type comment struct{}

// NewComment return the Comment service
func NewComment() Comment {
	return &comment{}
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
