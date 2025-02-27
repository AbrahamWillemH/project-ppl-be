package main

import (
	"log"
	"project-ppl-be/config"
	"project-ppl-be/server"
)

func main() {
	// Initialize the database
	config.ConnectDB()
	defer config.CloseDB()

	// Start the Gin server AFTER setting up the database
	router := server.SetupRouter()
	log.Fatal(router.Run(":8000"))
}
