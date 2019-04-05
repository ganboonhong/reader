package main

import (
	"log"
	"net/http"

	router "./router"
)

func main() {
	router.SetRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
