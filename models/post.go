package models

type Post struct {
	ID          int
	UserID      int
	Title       string
	CatergoryID []int
	Content     string
	Creat_at    string
	Update_to   string
}
