package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/grzany/versionski/config"
	"github.com/grzany/versionski/handler"
)

//Routes sets up all routes for the server
func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger, // Log API request calls
		//middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/common", handler.Routes(configuration))
		r.Mount("/prometheus", handler.PromRoutes(configuration))
	})

	return router
}
