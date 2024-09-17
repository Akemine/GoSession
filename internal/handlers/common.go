package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sand.com/config"
	model "sand.com/internal/models"
	"sand.com/internal/tools"
)

// protectedHandler gère une route protégée qui nécessite une authentification.
// Elle renvoie un message indiquant si l'utilisateur est connecté ou non.
// Est une route "Exemple" et inutile
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération de la session
	session, _ := config.Store.Get(r, "session-name")

	// Vérification de l'authentification
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Vous n'êtes pas connecté")
		return
	}

	// Utilisateur authentifié
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Vous êtes connecté")
}

// publicHandler gère une route publique accessible à tous les utilisateurs.
// Est une route "Exemple" et inutile
func PublicHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ceci est une route publique accessible à tous")
}

// profileHandler gère l'affichage du profil de l'utilisateur.
// Il récupère l'email de l'utilisateur à partir de la session et l'affiche.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "session-name")
	email := session.Values["email"].(string)

	db, err := tools.ConnectAndVerifyDb()
	if err != nil {
		http.Error(w, "Erreur de connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user model.UserProfile
	err = db.QueryRow("SELECT id, username, email, type, created_at FROM users WHERE email = $1", email).Scan(
		&user.ID, &user.Username, &user.Email, &user.Type, &user.CreatedAt)
	if err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
