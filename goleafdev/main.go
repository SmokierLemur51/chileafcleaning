package main

import (
	"net/http"
	"log"

	"github.com/SmokierLemur51/goleafdev/routes"

	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	PORT 	string
}


func main() {
	var config ServerConfig = ServerConfig{PORT: ":5000"}
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	routes.ConfigureRoutes(r)
	

	log.Println("Starting server on port ", config.PORT)
	http.ListenAndServe(config.PORT, r)
	// defer db.Close()
}