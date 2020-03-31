package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	tools "github.com/grzany/versionski/src/tools"
)

//Routes defines routes for common handler under /v1/api/common
func Routes(conf *tools.Config) *chi.Mux {
	router := chi.NewRouter()
	//router.Get("/{todoID}", GetATodo(configuration))
	//router.Delete("/{todoID}", DeleteTodo(configuration))
	//router.Post("/", CreateTodo(configuration))
	router.Get("/", GetDefaultRoute(conf))
	router.Get("/config", GetConfig(conf))
	return router
}

//GetDefaultRoute implements / route
func GetDefaultRoute(conf *tools.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Default route"
		render.JSON(w, r, response) // Return some demo response
	}
}

//GetConfig prints out config from file as a json
func GetConfig(conf *tools.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["Przywitanie"] = "Hello Adam, Hi there dude"
		render.JSON(w, r, response)
	}
}
