package main

import (
	"context"
	"encoding/json"
    "fmt"
	"log"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/barthr/newsapi"
)

type Article struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

func getNews() {
	c := newsapi.NewClient("07751a198b5440929cd22fc907b10389", newsapi.WithHTTPClient(http.DefaultClient))

	articles, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"techcrunch", "cnn", "time"},
	})

	if err != nil {
		panic(err)
	}

	// for _, s := range articles.Articles {
	//  fmt.Printf("%+v\n\n", s)
	// }

	b, err := json.Marshal(articles.Articles)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

func main() {
	// fmt.Printf("%+v", articles)

	// for _, a := range articles {
	//     fmt.Println(a.Title)
	// }

	// fmt.Println(string(b))

	http.HandleFunc("/list/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("tmpl/list.html")
		b, err := ioutil.ReadFile("debug.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(string(b))

		b, err = ioutil.ReadFile("debug.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		var Articles []Article
		err = json.Unmarshal(b, &Articles)

		t.Execute(w, Articles)
	})

    log.Fatal(http.ListenAndServe(":8080", nil))
}
