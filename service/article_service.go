package service

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	// "os"
	"strconv"
	"time"

	"github.com/ganboonhong/reader/model"

	"github.com/barthr/newsapi"
)

type ArticleService struct{}

const (
	PageSize = 10
	TypeEveryThing = "everything"
	TypeTopHeadline = "topheadline"
	NewsApiKey = "07751a198b5440929cd22fc907b10389"
)

func (a ArticleService) GetArticles(param *model.ArticlesParam) (*model.ArticleResult, error) {
	c := newsapi.NewClient(NewsApiKey, newsapi.WithHTTPClient(http.DefaultClient))
	var b []byte
	var ArticleResponse *newsapi.ArticleResponse
	var err error
	ctx := context.Background()

	switch param.NewsType {
	case TypeEveryThing:
		EverythingParameters := &newsapi.EverythingParameters{
			Sources: param.ArticleSources,
			From: param.DateStart,
			To: param.DateEnd,
			Page: param.Page,
			PageSize: PageSize,
			Language: "en",
		}
		ArticleResponse, err = c.GetEverything(ctx, EverythingParameters)

	case TypeTopHeadline:
		TopHeadlineParameters := &newsapi.TopHeadlineParameters{
			Page: param.Page,
			PageSize: PageSize,
			Country: param.Country,
		}
		ArticleResponse, err = c.GetTopHeadlines(ctx, TopHeadlineParameters)
	}

	b, err = json.Marshal(&ArticleResponse.Articles)

	var Articles []model.Article
	err = json.Unmarshal(b, &Articles)

	if err != nil{
		return nil, fmt.Errorf("could not fetch articles from server: %v", err)
	}

	ArticleResult := model.ArticleResult {
		TotalResults: ArticleResponse.TotalResults,
		Articles: Articles,
	}

	return &ArticleResult, nil
}

func (a ArticleService) GetStaticArticles(param *model.ArticlesParam)(*model.ArticleResult, error){
	b, err := ioutil.ReadFile("static/mock_data/article.json")
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
	// gopath := os.Getenv("GOPATH") // unit test will get "runtime error: invalid memory address or nil pointer dereference" if relative path is being used (go test ./... will break)
	// t, _ := template.ParseFiles(gopath + "/src/github.com/ganboonhong/reader/static/tmpl/article.html")
	// t, _ := template.ParseFiles("static/tmpl/article.html")
	t, _ := template.ParseFiles("/go/src/github.com/ganboonhong/reader/static/tmpl/article.html")
	var err error
	var i interface{}

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, i)
	// fmt.Fprint(w, "ok")
}

func (a ArticleService) GetArticleHandler(w http.ResponseWriter, r *http.Request){
	log.SetFlags(log.Lshortfile)
	ArticleService := ArticleService{}
	q := r.URL.Query()

	param, err := GetArticleParam(q)
	if err != nil {
		log.Println(err)
		return
	}

	// result, err := ArticleService.GetStaticArticles(param)
	result, err := ArticleService.GetArticles(param)
	if err != nil {
		log.Println(err)
		return
	}

	draw, err := strconv.Atoi(q["draw"][0])
	if err != nil {
		fmt.Println(err)
		return
	}

	data := model.DtResponse {
		Draw: draw,
		RecordsTotal: result.TotalResults,
		RecordsFiltered: result.TotalResults,
		Data: result.Articles,
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	 w.Header().Set("Content-Type", "application/json")
	 w.Write(jsonStr)
}

func GetArticleParam(q url.Values) (*model.ArticlesParam, error) {
	sd, err := time.Parse(time.RFC3339, q["s_date"][0] + "T00:00:00+08:00")
	if err != nil {
		return nil, err
	}
	ed, err := time.Parse(time.RFC3339, q["e_date"][0] + "T23:59:59+08:00")
	if err != nil {
		return nil, err
	}

	pg, err := strconv.Atoi(q["page"][0])
	if err != nil {
		return nil, err
	}
	pg += 1;

	p := &model.ArticlesParam{
		ArticleSources : q["article_sources[]"],
		Country : q["country"][0],
		DateEnd: ed,
		DateStart: sd,
		NewsType: q["news_type"][0],
		Page: pg,
	}

	return p, nil
}