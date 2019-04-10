package main

import (
	"log"
	"net/http"

	"github.com/ganboonhong/reader/router"
)

func main() {
	router.SetRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
