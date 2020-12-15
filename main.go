package main

import (
	"devbook/src/config"
	router "devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Load()

	router := router.GetRouter()

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port), router))

	fmt.Sprintf("Listening on port %d", config.Port)
}