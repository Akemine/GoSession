package middleware

import (
	"net/http"

	"sand.com/config"
	logger "sand.com/pkg"
)

// isAuthenticated est un middleware qui vérifie si l'utilisateur est authentifié.
// Il est utilisé pour protéger les routes qui nécessitent une authentification.
func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := config.Store.Get(r, "session-name")
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Error(w, "Vous n'êtes pas connecté", http.StatusUnauthorized)
			return
		}
		email := session.Values["email"].(string)
		logger.LogAPICall(r.URL.Path, r.Method, email, http.StatusOK)
		next.ServeHTTP(w, r)
	}
}
