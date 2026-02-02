package routes

import (
	"AP_Final/handlers"
	"net/http"
)

func RegisterRoutes() {
	// Публичные роуты
	http.HandleFunc("POST /register", handlers.Register)
	http.HandleFunc("POST /login", handlers.Login)
	http.HandleFunc("GET /home", handlers.Home)

	// Защищенные роуты (через AuthMiddleware)
	http.HandleFunc("GET /course", handlers.AuthMiddleware(handlers.GetAllCourses))
	http.HandleFunc("GET /course/{id}", handlers.AuthMiddleware(handlers.GetCourse))
	http.HandleFunc("PUT /course/{id}", handlers.AuthMiddleware(handlers.UpdateCourse))
	http.HandleFunc("DELETE /course/{id}", handlers.AuthMiddleware(handlers.DeleteCourse))

	// Статика
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
