package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xpdemon/userDb/appUser"
	"log"
	"os"
)

func Initialize() *sql.DB {

	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		err = os.MkdirAll("./data", 0755)
		if err != nil {
			log.Fatalf("Erreur lors de la création du répertoire : %v", err)
		}
	}

	// Connexion à la base SQLite
	db, err := sql.Open("sqlite3", "./data/users.db")
	if err != nil {
		log.Fatal(err)
	}
	// Création de la table "users"
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	_, err = Add(db, &appUser.AppUser{Username: "admin", Password: "secret"})
	if err != nil {
		fmt.Println("impossible de crée l'admin")
	}

	fmt.Println("Base de données initialisée et table créée.")
	return db
}
