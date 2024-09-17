package internal

import (
	"github.com/gorilla/mux"
	"sand.com/internal/handlers"
	"sand.com/internal/middleware"
)

// SetupRoutes setup the routes
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/protected", middleware.IsAuthenticated(handlers.ProtectedHandler)).Methods("GET")
	r.HandleFunc("/logout", middleware.IsAuthenticated(handlers.LogoutHandler)).Methods("POST")
	r.HandleFunc("/profile", middleware.IsAuthenticated(handlers.ProfileHandler)).Methods("GET")
	r.HandleFunc("/public", handlers.PublicHandler).Methods("GET")
}
