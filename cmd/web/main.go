package main

import (
	"flag"
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

	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// All URL paths that starts with "/static/"
	// Strip "/static" prefix before the request reachers the file server
	mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

	// Regiter the other applications routes as normal
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it. Note that we're using the
	// log.Printf() function to interpolate the address with the log message.
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)

	log.Fatal(err)
}
