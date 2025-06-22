// Package swagger Documentation
//
// @title Project PPL API
// @version 0.2.1
// @description API documentation for Project PPL - Learnify
// @schemes https
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

	utils "project-ppl-be/src/utils"

	"github.com/rs/cors"
)

func main() {
	// Initialize the database
	db := config.ConnectDB()
	defer db.Close()

	// Set up the Gin router with db
	router := server.SetupRouter()

	// Create a CORS wrapper with default settings
	corsHandler := cors.AllowAll()

	// Wrap the Gin router with CORS
	handler := corsHandler.Handler(router)

	// Start Cron
	utils.StartCron(db)

	// Start the server with the wrapped handler
	log.Fatal(http.ListenAndServe(":8080", handler))
}
