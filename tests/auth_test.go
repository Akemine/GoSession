package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"sand.com/config"
	"sand.com/internal/handlers"
)

// TestLoginHandler vérifie le bon fonctionnement du processus de connexion.
// Il simule une requête de connexion avec des identifiants valides et vérifie
// que le handler répond avec le statut et le message attendus.
func TestLoginHandler(t *testing.T) {
	// Création des données de test pour la connexion
	loginData := map[string]string{
		"email":    "EtLongMucle@gmail.com",
		"password": "ifone16pro",
	}
	body, _ := json.Marshal(loginData)

	// Création de la requête de test
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Création d'un ResponseRecorder pour enregistrer la réponse
	rr := httptest.NewRecorder()

	// Appel du handler
	handler := http.HandlerFunc(handlers.LoginHandler)
	handler.ServeHTTP(rr, req)

	// Vérification du code de statut
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler a retourné un mauvais code de statut : got %v want %v", status, http.StatusOK)
	}

	// Vérification du corps de la réponse
	expected := "Connexion réussie"
	if rr.Body.String() != expected {
		t.Errorf("Handler a retourné un corps inattendu : got %v want %v", rr.Body.String(), expected)
	}
}

// TestLogoutHandler vérifie le bon fonctionnement du processus de déconnexion.
// Il simule une session authentifiée, envoie une requête de déconnexion,
// et vérifie que le handler répond correctement en supprimant l'authentification.
func TestLogoutHandler(t *testing.T) {
	// Création de la requête de test
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Création d'un ResponseRecorder pour enregistrer la réponse
	rr := httptest.NewRecorder()

	// Simulation d'une session authentifiée
	session, _ := config.Store.Get(req, "session-name")
	session.Values["authenticated"] = true
	session.Values["email"] = "test@example.com"
	session.Save(req, rr)

	// Appel du handler
	handler := http.HandlerFunc(handlers.LogoutHandler)
	handler.ServeHTTP(rr, req)

	// Vérification du code de statut
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler a retourné un mauvais code de statut : got %v want %v", status, http.StatusOK)
	}

	// Vérification du corps de la réponse
	expected := "Déconnexion réussie"
	if rr.Body.String() != expected {
		t.Errorf("Handler a retourné un corps inattendu : got %v want %v", rr.Body.String(), expected)
	}
}
