package main

import (
	"log"
	"net/http"

	router "./router"
)

func main() {
	router.Routes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
