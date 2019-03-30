package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./service"
)

func main() {

	http.HandleFunc("/list3/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("tmpl/list3.html")
		ArticleService := service.ArticleService{}
		articles, err := ArticleService.GetArticles()
		// articles, err := ArticleService.GetStaticArticles()

		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, articles)
	})

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("tmpl/style"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
