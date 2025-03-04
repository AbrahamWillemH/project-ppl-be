// Package swagger Documentation
//
// @title Project PPL API
// @version 0.1.3
// @description API documentation for Project PPL - Kuda Hitam
// @schemes http
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer {token}" to authenticate
package main

import (
	"log"
	"net/http"
	"project-ppl-be/config"
	_ "project-ppl-be/docs"
	"project-ppl-be/src/server"

	"github.com/rs/cors"
)

func main() {
	// Initialize the database
	config.ConnectDB()
	defer config.CloseDB()

	// Set up the Gin router
	router := server.SetupRouter()

	// Create a CORS wrapper with default settings
	corsHandler := cors.New(cors.Options{
	    AllowedOrigins:   []string{"http://localhost:5173"}, // Allow only your frontend
	    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	    AllowedHeaders:   []string{"Authorization", "Content-Type"},
	    AllowCredentials: true,
	})

	// Wrap the Gin router with CORS
	handler := corsHandler.Handler(router)

	// Start the server with the wrapped handler
	log.Fatal(http.ListenAndServe(":8080", handler))
}
