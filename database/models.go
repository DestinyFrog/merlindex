package database

type Model struct {
	Id int
}

type User struct {
	Model
	Name     string
	Email    string
	Password string
}

type List struct {
	Model
	Title  string
	UserId int
}

type ListItem struct {
	Model
	Title  string
	UserId int
	ListId int
}

type Comment struct {
	Model
	Message string
}

type CommentListItem struct {
	Model
	CommentId  int
	ListItemId int
}
