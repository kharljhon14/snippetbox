package main

import (
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

func (app *application) routes() http.Handler {
	// Router
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// All URL paths that starts with "/static/"
	// Strip "/static" prefix before the request reachers the file server
	mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.logRequest(secureHeaders(mux))
}
