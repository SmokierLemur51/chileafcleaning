package routes

import (
	"github.com/SmokierLemur51/goleafdev/handlers"

	"github.com/go-chi/chi/v5"
)

type EndPoints struct {
	Index 		string
	About 		string
	Services    string
}



func ConfigureRoutes(router *chi.Mux) {
	router.Get("/", handlers.IndexHandler)
	// router.Get("/services", handlers.ServiceHandler)
	// router.Get("/about", handlers.AboutHandler)

	router.Get("/em-application/{target}-{subject}-{message}", handlers.SendEmails)
	// https://go-chi.io/#/pages/routing
}

