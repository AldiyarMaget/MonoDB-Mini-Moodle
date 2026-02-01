package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"AP_Final/db"
	"AP_Final/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	db.Connect(mongoURI)

	routes.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is running on :" + port)
	http.ListenAndServe(":"+port, nil)
}
