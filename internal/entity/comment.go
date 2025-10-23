package entity



type Comment struct {
	ID uint
	PostID uint
	AuthorID uint
	Content string
}