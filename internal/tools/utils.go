package tools

import (
	"database/sql"

	"sand.com/internal/databases"
)

// ConnectAndVerifyDb établit une connexion à la base de données et vérifie qu'elle est valide.
// Elle retourne un pointeur vers la connexion SQL et une erreur si la connexion échoue.
func ConnectAndVerifyDb() (*sql.DB, error) {
	db, err := databases.ConnectDb()
	if err != nil {
		return nil, err
	}
	return db, nil
}
