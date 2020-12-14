package main

import (
	router "devbook/src/router"
	"log"
	"net/http"
)

func main() {

	router := router.GetRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}