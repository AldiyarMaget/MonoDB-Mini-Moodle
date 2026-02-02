package routes

import (
	"net/http"

	"AP_Final/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("GET /course/{id}", handlers.GetCourse)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("GET /home", handlers.Home)
}
