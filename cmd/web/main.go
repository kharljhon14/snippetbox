package main

import (
	"log"
	"net/http"
	"strings"
)

func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Router
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// All URL paths that starts with "/static/"
	// Strip "/static" prefix before the request reachers the file server
	mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

	// Regiter the other applications routes as normal
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on : 8000")
	err := http.ListenAndServe(":8000", mux)

	log.Fatal(err)
}
