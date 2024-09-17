package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// connectDb établit une connexion à la base de données PostgreSQL.
// Elle retourne une erreur si la connexion échoue.
func ConnectDb() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la connexion à la base de données : %v", err)
	}

	// Vérifier que la connexion est bien établie
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erreur lors du ping de la base de données : %v", err)
	}

	fmt.Println("Connexion à la base de données établie avec succès")
	return db, nil
}
