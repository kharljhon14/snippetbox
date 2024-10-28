package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Controllers
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// Check if the current URL path exactly matches "/"
	// If it doesn't use the http.NotFound() to send a 404

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	file := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(file...)

	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
		return
	}

	// err = ts.Execute(w, nil)

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.errorLog.Panicln(err.Error())
		app.serverError(w, err)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d ...", id)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// Set will overwrite if the header already exists
	w.Header().Set("Cache-Control", "public max-age-31536000")
	w.Header().Set("Content-Type", "application/json")
	w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
	w.Header()["Date"] = nil
	// In contrast, the Add() method appends a new "Cache-Control" header and can
	// be called multiple times.
	//* w.Header().Add("Cache-Control", "public")
	//* w.Header().Add("Cache-Control", "max-age=31536000")

	// The r.Method will check if the request type is using POST
	if r.Method != http.MethodPost {
		// w.WriteHeader() method will send a 405 status
		// w.Write() method to write a "Method Not Allowed"

		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte(`{"name": "Kharl"}`))
}
