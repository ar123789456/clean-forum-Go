package domain

type Comment struct {
	ID       int
	UserID   int
	PostId   int
	UserName string
	Text     string
	Creat_at string
}
