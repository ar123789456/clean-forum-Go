package domain

type Post struct {
	ID          int
	UserID      int
	Title       string
	CatergoryID int
	Tags        []Tag
	Content     string
	Creat_at    string
	Update_to   string
}
