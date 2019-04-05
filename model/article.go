package model

type Article struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
	Source  	Source `json:source`
}

type Source struct {
	Id string `json:"id"`
	Name string `json:"name"`
}