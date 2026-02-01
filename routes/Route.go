package routes

import (
	"net/http"

	"AP_Final/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
}
