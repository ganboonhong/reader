// the reason why creating new structs instead of using the structs in github.com/barthr/newsapi is:
// we can add some custom fields without changing code in github.com/barthr/newsapi

package model

import "time"

type Article struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
	Source  	Source `json:source`
	TotalResults int `json:totalResults`
}

type Source struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type ArticleResult struct {
	TotalResults int `json:"totalResults"`
	Articles []Article `json:"articles"`
}

type ArticlesParam struct {
	ArticleSources []string
	Country string
	DateStart time.Time
	DateEnd time.Time
	NewsType string
	Page int
}

type DtResponse struct {
	Draw int `json:"draw"`
	RecordsTotal int `json:"recordsTotal"`
	RecordsFiltered int `json:"recordsFiltered"`
	Data []Article `json:"data"`
}