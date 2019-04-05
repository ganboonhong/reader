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
	Draw int `json:"draw"`
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

func (a ArticleService) GetArticles(param GetArticlesParam) (model.ArticleResult, error) {
	c := newsapi.NewClient("07751a198b5440929cd22fc907b10389", newsapi.WithHTTPClient(http.DefaultClient))

	// articles, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		// Sources: []string{"techcrunch", "cnn", "time"},
	// })

	result, err := c.GetEverything(context.Background(), &newsapi.EverythingParameters{
		// Sources: []string{"techcrunch", "cnn", "time"},
		Sources: param.article_sources,
		From: param.s_date,
		To: param.e_date,
		Page: param.page,
		PageSize: 10,
		Language: "en",
		SortBy: "popularity",
	})

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(result.Articles)

	var Articles []model.Article
	err = json.Unmarshal(b, &Articles)

	if err != nil{
		return model.ArticleResult{}, fmt.Errorf("could not fetch articles from server: %v", err)
	}

	ArticleResult := model.ArticleResult {
		TotalResults: result.TotalResults,
		Articles: Articles,
	}

	return ArticleResult, nil
}

func (a ArticleService) GetStaticArticles()(*model.ArticleResult, error){
	b, err := ioutil.ReadFile("debug.json")
	if err != nil {
		return nil, fmt.Errorf("could not fetch static articles: %v", err)
	}

	var ArticleResult model.ArticleResult
	err = json.Unmarshal(b, &ArticleResult)
	if err != nil {
		return nil, fmt.Errorf("could not decode json: %v", err)
	}

	return &ArticleResult, nil
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
	
	page, _ := strconv.Atoi(query["page"][0])
	page += 1;
	
	param := GetArticlesParam{
		s_date : s_date,
		e_date : e_date,
		page : page,
		article_sources : query["article_sources[]"],
	}
	result, err := ArticleService.GetArticles(param)
	
	// result, err := ArticleService.GetStaticArticles()

	if err != nil {
		fmt.Println(err)
		return
	}

	draw, err := strconv.Atoi(query["draw"][0])

	if err != nil {
		fmt.Println(err)
		return
	}

	data := ArticleData {
		Draw: draw,
		RecordsTotal: result.TotalResults,
		RecordsFiltered: result.TotalResults,
		Data: result.Articles,
	}

	jsonStr, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	 w.Header().Set("Content-Type", "application/json")
	 w.Write(jsonStr)
}