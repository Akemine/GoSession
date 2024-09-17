package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"sand.com/config"
	model "sand.com/internal/models"
	"sand.com/internal/tools"
	logger "sand.com/pkg"
)

// loginHandler gère l'authentification des utilisateurs.
// Il vérifie les identifiants fournis et crée une session si l'authentification réussit.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Vérification si l'utilisateur est déjà connecté
	session, _ := config.Store.Get(r, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		logger.LogAPICall(r.URL.Path, r.Method, session.Values["email"].(string), http.StatusOK)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Vous êtes déjà connecté")
		return
	}

	var creds model.Credentials

	// Décodage du JSON
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := tools.ConnectAndVerifyDb()
	// Vérification des identifiants
	var hashedPassword string
	err = db.QueryRow("SELECT password_hash FROM users WHERE email = $1", creds.Email).Scan(&hashedPassword)
	if err != nil {
		http.Error(w, "Identifiants invalides", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Identifiants invalides", http.StatusUnauthorized)
		return
	}

	// Création de la session
	session.Values["authenticated"] = true
	session.Values["email"] = creds.Email
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
		return
	}

	logger.LogAPICall(r.URL.Path, r.Method, creds.Email, http.StatusOK)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Connexion réussie")
}

// logoutHandler gère la déconnexion des utilisateurs.
// Il invalide la session actuelle de l'utilisateur.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération de la session
	session, _ := config.Store.Get(r, "session-name")

	// Vérification si l'utilisateur est actuellement connecté
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		logger.LogAPICall(r.URL.Path, r.Method, "anonymous", http.StatusBadRequest)
		http.Error(w, "Vous n'êtes pas connecté", http.StatusBadRequest)
		return
	}

	// Récupération de l'email avant de réinitialiser la session
	email := session.Values["email"].(string)

	// Réinitialisation des valeurs de la session
	session.Values["authenticated"] = false
	session.Values["email"] = ""

	// Sauvegarde des modifications
	err := session.Save(r, w)
	if err != nil {
		logger.LogAPICall(r.URL.Path, r.Method, email, http.StatusInternalServerError)
		http.Error(w, "Erreur lors de la déconnexion", http.StatusInternalServerError)
		return
	}

	logger.LogAPICall(r.URL.Path, r.Method, email, http.StatusOK)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Déconnexion réussie")
}
