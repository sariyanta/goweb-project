package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sariyanta/goweb/pkg/config"
	"github.com/sariyanta/goweb/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	// Mux stands for 'HTTP request multiplexer'
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fs := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fs))

	return mux
}
