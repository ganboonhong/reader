package service

import (
	// "fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/ganboonhong/reader/model"
)

func TestGetArticleParam(t *testing.T) {
	u, err := url.Parse("http://localhost:8080/get_article?draw=2&s_date=2019-04-05&e_date=2019-04-05&article_sources%5B%5D=cnn&country=tw&page=0&dt%5Bdraw%5D=1&dt%5Bcolumns%5D%5B0%5D%5Bdata%5D=title&dt%5Bcolumns%5D%5B0%5D%5Bname%5D=&dt%5Bcolumns%5D%5B0%5D%5Bsearchable%5D=true&dt%5Bcolumns%5D%5B0%5D%5Borderable%5D=false&dt%5Bcolumns%5D%5B0%5D%5Bsearch%5D%5Bvalue%5D=&dt%5Bcolumns%5D%5B0%5D%5Bsearch%5D%5Bregex%5D=false&dt%5Bcolumns%5D%5B1%5D%5Bdata%5D=description&dt%5Bcolumns%5D%5B1%5D%5Bname%5D=&dt%5Bcolumns%5D%5B1%5D%5Bsearchable%5D=true&dt%5Bcolumns%5D%5B1%5D%5Borderable%5D=false&dt%5Bcolumns%5D%5B1%5D%5Bsearch%5D%5Bvalue%5D=&dt%5Bcolumns%5D%5B1%5D%5Bsearch%5D%5Bregex%5D=false&dt%5Bcolumns%5D%5B2%5D%5Bdata%5D=publishedAt&dt%5Bcolumns%5D%5B2%5D%5Bname%5D=&dt%5Bcolumns%5D%5B2%5D%5Bsearchable%5D=true&dt%5Bcolumns%5D%5B2%5D%5Borderable%5D=false&dt%5Bcolumns%5D%5B2%5D%5Bsearch%5D%5Bvalue%5D=&dt%5Bcolumns%5D%5B2%5D%5Bsearch%5D%5Bregex%5D=false&dt%5Bcolumns%5D%5B3%5D%5Bdata%5D=url&dt%5Bcolumns%5D%5B3%5D%5Bname%5D=&dt%5Bcolumns%5D%5B3%5D%5Bsearchable%5D=true&dt%5Bcolumns%5D%5B3%5D%5Borderable%5D=false&dt%5Bcolumns%5D%5B3%5D%5Bsearch%5D%5Bvalue%5D=&dt%5Bcolumns%5D%5B3%5D%5Bsearch%5D%5Bregex%5D=false&dt%5Bstart%5D=0&dt%5Blength%5D=10&dt%5Bsearch%5D%5Bvalue%5D=&dt%5Bsearch%5D%5Bregex%5D=false&news_type=topheadline&_=1554855422749")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	sd, err := time.Parse(time.RFC3339, q["s_date"][0]+"T00:00:00+08:00")
	if err != nil {
		t.Errorf("time.Parse error: %v", err)
	}
	ed, err := time.Parse(time.RFC3339, q["e_date"][0]+"T23:59:59+08:00")
	if err != nil {
		t.Errorf("time.Parse error: %v", err)
	}

	pg, err := strconv.Atoi(q["page"][0])
	if err != nil {
		t.Errorf("strconv.Atoi error: %v", err)
	}
	pg += 1

	want := &model.ArticlesParam{
		ArticleSources: q["article_sources[]"],
		Country:        q["country"][0],
		DateEnd:        ed,
		DateStart:      sd,
		NewsType:       q["news_type"][0],
		Page:           pg,
	}

	got, err := GetArticleParam(q)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetArticleParam: got %v, want %v", got, want)
	}
}

func TestArticlePageHandler(t *testing.T) {

	ArticleService := ArticleService{}
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(ArticleService.ArticlePageHandler)
	handler.ServeHTTP(w, r)
	resp := w.Result()

	// if resp.StatusCode != http.StatusOK {
	if resp.StatusCode != 300 {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
	}

}
