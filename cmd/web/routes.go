package main

import (
	"net/http"

	"github.com/yikuanzz/snippet/ui"
)

func (app *application) routes() *http.ServeMux {
	// Create http router
	mux := http.NewServeMux()

	// Register file server
	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("/static/", fileServer)

	// Register routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
