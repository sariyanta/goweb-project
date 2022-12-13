package main

import (
	"net/http"

	"github.com/sariyanta/goweb-project/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)

	// Create a file server which serves files out of the "./static/" directory.
	fs := http.FileServer(http.Dir("./static/"))

	// Tell the server to handle all requests starting with "/static/" by using
	// the file server we just created.
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	_ = http.ListenAndServe(":8080", nil)
}
