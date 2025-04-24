package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/yikuanzz/snippet/ui"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"html/base.tmpl",
		"html/partials/nav.tmpl",
		"html/pages/home.tmpl",
	}

	ts, err := template.ParseFS(ui.Templates, files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	_, err := w.Write([]byte("Create a new snippet..."))
	if err != nil {
		app.serverError(w, err)
	}
}
