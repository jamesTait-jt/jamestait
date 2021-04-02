package main

import (
	"fmt"
	"log"
	"net/http"
)

var PORT = 3000

func main() {
	http.HandleFunc("/", homePageHandler)
	fmt.Printf("Server listening on port %d\n", PORT)
	log.Panic(
		http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil),
	)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}