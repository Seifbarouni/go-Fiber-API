package models

type Article struct {
	ID       int    `json:id`
	AuthorId int    `json:author_id`
	Title    string `json:title`
	Content  string `json:content`
}
