package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"../service"
)

type Route struct {
	Method string
	Path string
	HandlerFunc http.Handler
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func SetRouter(){
	var router = mux.NewRouter().StrictSlash(true)
	ArticleService := service.ArticleService{}

	router.HandleFunc("/article", ArticleService.ArticlePageHandler)
	router.HandleFunc("/get_article", ArticleService.GetArticleHandler)
	router.HandleFunc("/favicon.ico", faviconHandler)
	
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("static/style"))))
	http.Handle("/bower_components/", http.StripPrefix("/bower_components/", http.FileServer(http.Dir("static/bower_components"))))
	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("static/node_modules"))))
	http.Handle("/", router)
}