package models

type Post struct {
	Id 		string
	Content string
	Hash	string
}

func NewPost(id, content, hash string) *Post {
	return &Post{id, content, hash}
}