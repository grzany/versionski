package versionski

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	handlers "github.com/grzany/versionski/src/handlers"
	tools "github.com/grzany/versionski/src/tools"
)

//Routes sets up all routes for the server
func Routes(configuration *tools.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger, // Log API request calls
		//middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/common", handlers.Routes(configuration))
	})

	return router
}
