package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"sand.com/internal"
	"sand.com/internal/databases"
	logger "sand.com/pkg"
)

func main() {
	// Configuration du logger
	if err := logger.Init(); err != nil {
		log.Fatalf("Erreur lors de l'initialisation du logger : %v", err)
	}
	defer logger.Close()

	// Connexion à la base de données
	db, err := databases.ConnectDb()
	if err != nil {
		logger.Logger.Fatalf("Erreur lors de la connexion à la base de données : %v", err)
	}
	defer db.Close()

	// Configuration du routeur
	r := mux.NewRouter()
	internal.SetupRoutes(r)

	// Configuration des CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Création du serveur HTTP
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      c.Handler(r),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Démarrage du serveur dans une goroutine
	go func() {
		logger.Logger.Println("Serveur démarré sur http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatalf("Erreur lors du démarrage du serveur : %v", err)
		}
	}()

	// Configuration de la gestion de l'arrêt gracieux
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	// Blocage jusqu'à ce qu'un signal soit reçu
	<-channel

	// Création d'un contexte avec un délai pour l'arrêt gracieux
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Arrêt gracieux du serveur
	srv.Shutdown(ctx)

	logger.Logger.Println("Arrêt du serveur")
	os.Exit(0)
}
