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
	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // дефолт, если .env пуст
	}

	db.Connect(mongoURI)

	// Инициализируем маршруты
	routes.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Сервер запущен на http://localhost:%s/home", port)
	// nil означает использование стандартного ServeMux, в который мы пишем через http.HandleFunc
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
