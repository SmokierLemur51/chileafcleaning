package main

import (
	"net/http"
	"log"

	"github.com/SmokierLemur51/chileafcleaning/config"
	"github.com/SmokierLemur51/chileafcleaning/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func configureServer() config.ServerConfig {
	return config.ServerConfig{
		Port: ":" + config.PORT,
	}
}

func main(){
	SERVER_CONFIG := configureServer()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	routes.ConfigureRoutes(router)

	http.ListenAndServe(SERVER_CONFIG.Port, router)
	log.Println("Started server on port ", SERVER_CONFIG.Port)
}
