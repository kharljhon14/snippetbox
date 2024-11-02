package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kharljhon14/snippetbox/internal/models"
)

// Controllers
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// Check if the current URL path exactly matches "/"
	// If it doesn't use the http.NotFound() to send a 404

	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }
	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.tmpl.html", data)

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, the values of any named parameters
	// will be stored in the request context.

	params := httprouter.ParamsFromContext(r.Context())

	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}

		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.tmpl.html", data)

}

// func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

// Set will overwrite if the header already exists
// w.Header().Set("Cache-Control", "public max-age-31536000")
// w.Header().Set("Content-Type", "application/json")
// w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}

// w.Header()["Date"] = nil
// In contrast, the Add() method appends a new "Cache-Control" header and can
// be called multiple times.
//* w.Header().Add("Cache-Control", "public")
//* w.Header().Add("Cache-Control", "max-age=31536000")

// The r.Method will check if the request type is using POST
// if r.Method != http.MethodPost {
// 	// w.WriteHeader() method will send a 405 status
// 	// w.Write() method to write a "Method Not Allowed"

// 	w.Header().Set("Allow", http.MethodPost)
// 	// w.WriteHeader(405)
// 	// w.Write([]byte("Method Not Allowed"))

// 	app.clientError(w, http.StatusMethodNotAllowed)
// 	return
// }

// title := "0 Snail"
// 	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
// 	expires := 7

// 	id, err := app.snippets.Insert(title, content, expires)

// 	if err != nil {
// 		app.serverError(w, err)
// 	}

// 	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
// }

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
