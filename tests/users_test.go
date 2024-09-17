package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"sand.com/internal/handlers"
)

// Username must be unique
func TestRegisterHandler(t *testing.T) {

	// Création des données de test
	userData := map[string]interface{}{
		"username":   "TestUser2",
		"email":      "testuser2@example.com",
		"password":   "$2a$10$o/VD7r5Zm6oDjLOc2vLd.OZKCQdrgKaI8X3eCfRjPycs5c6yOF8c62",
		"type":       true,
		"created_at": time.Now(),
	}
	body, _ := json.Marshal(userData)

	// Création de la requête de test
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Création d'un ResponseRecorder pour enregistrer la réponse
	rr := httptest.NewRecorder()

	// Appel du handler
	handler := http.HandlerFunc(handlers.RegisterHandler)
	handler.ServeHTTP(rr, req)

	// Vérification du code de statut
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler a retourné un mauvais code de statut : got %v want %v", status, http.StatusCreated)
	}

	// Vérification du corps de la réponse
	expected := "Utilisateur enregistré avec succès"
	if rr.Body.String() != expected {
		t.Errorf("Handler a retourné un corps inattendu : got %v want %v", rr.Body.String(), expected)
	}
}
