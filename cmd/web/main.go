package main

import (
	"log"
	"net/http"
)

func main() {
	// Router
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on : 8000")
	err := http.ListenAndServe(":8000", mux)

	log.Fatal(err)
}
