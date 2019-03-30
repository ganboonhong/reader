package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"html/template"
	"net/http"

	"../model"

	"github.com/barthr/newsapi"
)

type ArticleService struct{

}

func (a ArticleService) GetArticles() ([]model.Article, error) {
	c := newsapi.NewClient("07751a198b5440929cd22fc907b10389", newsapi.WithHTTPClient(http.DefaultClient))

	articles, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"techcrunch", "cnn", "time"},
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

func (a ArticleService) ArticleHandler(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("tmpl/list3.html")
	ArticleService := ArticleService{}
	// articles, err := ArticleService.GetArticles()
	articles, err := ArticleService.GetStaticArticles()

	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, articles)
}