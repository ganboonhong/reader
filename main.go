package main

import (
	"log"
	"net/http"

	"github.com/ganboonhong/reader/router"
)

const (
	port = "8080"
)

func main() {
	router.SetRouter()
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
