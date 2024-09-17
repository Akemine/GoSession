package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	model "sand.com/internal/models"
	"sand.com/internal/tools"
)

// registerHandler gère l'enregistrement de nouveaux utilisateurs.
// Il valide les données d'entrée, hache le mot de passe et enregistre l'utilisateur dans la base de données.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Décodage du JSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hachage du mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
		return
	}

	db, err := tools.ConnectAndVerifyDb()
	// Insertion dans la base de données
	_, err = db.Exec(`
		INSERT INTO users (username, email, password_hash, type, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`, user.Username, user.Email, string(hashedPassword), user.Type, time.Now())

	if err != nil {
		http.Error(w, "Erreur lors de l'insertion dans la base de données", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Utilisateur enregistré avec succès")
}
