package main

import (
	"log"
	"net/http"

	"ticketing-system/config"
	"ticketing-system/middleware" // 👈 import it
	"ticketing-system/routes"
)

func main() {
	config.ConnectDB()

	r := routes.RegisterRoutes()

	// 🪵 Attach logger middleware
	loggedRouter := middleware.RequestLogger(r)

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
