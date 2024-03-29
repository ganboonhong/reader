package router

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/ganboonhong/reader/service"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func SetRouter() {
	var router = mux.NewRouter().StrictSlash(true)
	ArticleService := service.ArticleService{}

	router.HandleFunc("/", ArticleService.ArticlePageHandler)
	router.HandleFunc("/get_article", ArticleService.GetArticleHandler)
	router.HandleFunc("/favicon.ico", faviconHandler)

	gopath := os.Getenv("GOPATH")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(gopath+"/src/github.com/ganboonhong/reader/static/"))))
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.Handle("/", router)
}
