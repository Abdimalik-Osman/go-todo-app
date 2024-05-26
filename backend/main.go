package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-react-todo/router"
)

func main() {
	router := router.Router()
	fmt.Println("Starting server at port: 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
