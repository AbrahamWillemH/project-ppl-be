package main

import (
	"log"

	"github.com/gin-contrib/cors"

	"project-ppl-be/config"
	_ "project-ppl-be/docs"
	"project-ppl-be/server"
)

func main() {
	// Initialize the database
	config.ConnectDB()
	defer config.CloseDB()

	// Start the Gin server AFTER setting up the database
	router := server.SetupRouter()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	log.Fatal(router.Run(":8000"))
}
