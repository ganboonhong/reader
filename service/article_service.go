package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"html/template"
	"net/http"
	// "reflect"
	"strconv"
	"time"

	"../model"

	"github.com/barthr/newsapi"
)

type ArticleService struct{

}

type ArticleData struct {
	Draw int64 `json:"draw"`
	RecordsTotal int `json:"recordsTotal"`
	RecordsFiltered int `json:"recordsFiltered"`
	Data []model.Article `json:"data"`
}

type GetArticlesParam struct {
	article_sources []string
	s_date time.Time
	e_date time.Time
	page int
}

type Source struct {
	source_name string `json:"name"`
}

func (a ArticleService) GetArticles(param GetArticlesParam) ([]model.Article, error) {
	c := newsapi.NewClient("07751a198b5440929cd22fc907b10389", newsapi.WithHTTPClient(http.DefaultClient))

	// articles, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		// Sources: []string{"techcrunch", "cnn", "time"},
	// })

	articles, err := c.GetEverything(context.Background(), &newsapi.EverythingParameters{
		// Sources: []string{"techcrunch", "cnn", "time"},
		Sources: param.article_sources,
		From: param.s_date,
		To: param.e_date,
		Page: param.page,
		Language: "en",
		SortBy: "popularity",
	})

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(articles.Articles)

	var Articles []model.Article
	err = json.Unmarshal(b, &Articles)

	if err != nil{
		return nil, fmt.Errorf("could not fetch articles from server: %v", err)
	}

	return Articles, nil
}

func (a ArticleService) GetStaticArticles()([]model.Article, error){
	b, err := ioutil.ReadFile("debug.json")
	if err != nil {
		return nil, fmt.Errorf("could not fetch static articles: %v", err)
	}

	var Articles []model.Article
	err = json.Unmarshal(b, &Articles)
	if err != nil {
		return nil, fmt.Errorf("could not decode json: %v", err)
	}
	return Articles, nil
}

func (a ArticleService) ArticlePageHandler(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("static/tmpl/article.html")
	var err error
	var i interface{}

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, i)
}

func (a ArticleService) GetArticleHandler(w http.ResponseWriter, r *http.Request){
	ArticleService := ArticleService{}

	query := r.URL.Query()
	
	s_date, err := time.Parse(time.RFC3339, query["e_date"][0] + "T00:00:00Z")
	e_date, err := time.Parse(time.RFC3339, query["e_date"][0] + "T00:00:00Z")
	page, _ := strconv.Atoi(query["dt[start]"][0])
	page += 1
	param := GetArticlesParam{
		s_date : s_date,
		e_date : e_date,
		page : page,
		article_sources : query["article_sources[]"],
	}
	articles, err := ArticleService.GetArticles(param)
	
	// articles, err := ArticleService.GetStaticArticles()

	if err != nil {
		fmt.Println(err)
		return
	}

	draw, err := strconv.ParseInt(query["draw"][0], 10, 64)
	data := ArticleData {
		Draw: draw,
		RecordsTotal: len(articles),
		RecordsFiltered: len(articles),
		Data: articles,
	}

	jsonStr, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	 w.Header().Set("Content-Type", "application/json")
	 w.Write(jsonStr)
}