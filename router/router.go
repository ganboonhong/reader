package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"../service"
)

func Routes(){
	var router = mux.NewRouter()

	ArticleService := service.ArticleService{}
	router.HandleFunc("/list3/", ArticleService.ArticleHandler)

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("tmpl/style"))))

	http.Handle("/", router)
}