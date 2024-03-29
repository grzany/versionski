package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grzany/versionski/app"
	"github.com/grzany/versionski/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	configuration, err := config.NewConfig()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	router := app.Routes(configuration)
	router.Mount("/metrics", promhttp.Handler())

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Println("Serving application at PORT :" + configuration.Constants.PORT)
	log.Fatal(http.ListenAndServe(":"+configuration.Constants.PORT, router)) //
	// Serve metrics.
	log.Printf("serving metrics at: %s", ":9090")
	go http.ListenAndServe(":9090", promhttp.Handler())
}
