package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/grzany/versionski/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

//Routes defines routes for common handler under /v1/api/common
func Routes(conf *config.Config) *chi.Mux {
	router := chi.NewRouter()
	//router.Get("/{todoID}", GetATodo(configuration))
	//router.Delete("/{todoID}", DeleteTodo(configuration))
	//router.Post("/", CreateTodo(configuration))
	router.Get("/", GetDefaultRoute(conf))
	router.Get("/config", GetConfig(conf))
	return router
}

func PromRoutes(conf *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/ops", RecordMetrics())
	return router
}

func RecordMetrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}
}

//GetDefaultRoute implements / route
func GetDefaultRoute(conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Default route"
		render.JSON(w, r, response) // Return some demo response
	}
}

//GetConfig prints out config from file as a json
func GetConfig(conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["Przywitanie"] = "Hello Adam, Hi there dude"
		render.JSON(w, r, response)
	}
}
