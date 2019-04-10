package service

import (
	"log"
	"testing"
	"net/url"
)

func TestGetArticleParam(t *testing.T){
	u, err := url.Parse("http://localhost:8080/get_article?draw=2&s_date=2019-04-05&e_date=2019-04-05&article_sources%5B%5D=cnn&country=tw&page=0&dt%5Bdraw%5D=1&dt%5Bcolumns%5D%5B0%5D%5Bdata%5D=title&dt%5Bcolumns%5D%5B0%5D%5Bname%5D=&dt%5Bcolumns%5D%5B0%5D%5Bsearchable%5D=true&dt%5Bcolumns%5D%5B0%5D%5Borderable%5D=false&dt%5Bcolumns%5D%5B0%5D%5Bsearch%5D%5Bvalue%5D=&dt%5Bcolumns%5D%5B0%5D%5Bsearch%5D%5Bregex%5D=false&dt%5Bcolumns%5D%5B1%5D%5Bdata%5D=description&dt%5Bcolumns%5D%5B1%5D%5Bname%5D=&dt%5Bcolumns%5D%5B1%5D%5Bsearchable%5D=true&dt%5Bcolumns%5D%5B1%5D%5Borderable%5D=false&dt%5Bcolumns%5D%5B1%5D%5Bsearch%5D%5Bvalue%5D=&dt%5Bcolumns%5D%5B1%5D%5Bsearch%5D%5Bregex%5D=false&dt%5Bcolumns%5D%5B2%5D%5Bdata%5D=publishedAt&dt%5Bcolumns%5D%5B2%5D%5Bname%5D=&dt%5Bcolumns%5D%5B2%5D%5Bsearchable%5D=true&dt%5Bcolumns%5D%5B2%5D%5Borderable%5D=false&dt%5Bcolumns%5D%5B2%5D%5Bsearch%5D%5Bvalue%5D=&dt%5Bcolumns%5D%5B2%5D%5Bsearch%5D%5Bregex%5D=false&dt%5Bcolumns%5D%5B3%5D%5Bdata%5D=url&dt%5Bcolumns%5D%5B3%5D%5Bname%5D=&dt%5Bcolumns%5D%5B3%5D%5Bsearchable%5D=true&dt%5Bcolumns%5D%5B3%5D%5Borderable%5D=false&dt%5Bcolumns%5D%5B3%5D%5Bsearch%5D%5Bvalue%5D=&dt%5Bcolumns%5D%5B3%5D%5Bsearch%5D%5Bregex%5D=false&dt%5Bstart%5D=0&dt%5Blength%5D=10&dt%5Bsearch%5D%5Bvalue%5D=&dt%5Bsearch%5D%5Bregex%5D=false&news_type=topheadline&_=1554855422749")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	got := q["page"][0]
	want := "0"
	if got != want {
		t.Errorf("page: got %v, want %v", got, want)
	}

	got = q["s_date"][0]
	want = "2019-04-05"
	if got != want {
		t.Errorf("s_date: got %v, want %v", got, want)
	}

	got = q["e_date"][0]
	want = "2019-04-05"
	if got != want {
		t.Errorf("e_date: got %v, want %v", got, want)
	}

	got_article_sources := q["article_sources[]"]
	want_article_sources := []string{"cnn"}
	if len(got_article_sources) != len(want_article_sources) {
		t.Errorf("article_sources: got %v, want %v", got_article_sources, want_article_sources)
	}

	for i, v := range got_article_sources {
		if v != want_article_sources[i] {
			t.Errorf("article_sources: got %v, want %v", got_article_sources, want_article_sources)
		}
	}

	got = q["country"][0]
	want = "tw"
	if got != want {
		t.Errorf("country: got %v, want %v", got, want)
	}

	p, err := GetArticleParam(q)
	log.Println(p)
}