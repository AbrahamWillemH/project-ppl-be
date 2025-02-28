package main

import (
	"log"
	"net/http"
	"project-ppl-be/config"
	_ "project-ppl-be/docs"
	"project-ppl-be/server"

	"github.com/rs/cors"
)

func main() {
	// Initialize the database
	config.ConnectDB()
	defer config.CloseDB()

	// Set up the Gin router
	router := server.SetupRouter()

	// Create a CORS wrapper with default settings
	corsHandler := cors.Default()

	// Wrap the Gin router with CORS
	handler := corsHandler.Handler(router)

	// Start the server with the wrapped handler
	log.Fatal(http.ListenAndServe(":8000", handler))
}
