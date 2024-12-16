package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/xpdemon/userDb/appUser"
	"golang.org/x/crypto/bcrypt"
)

func Add(db *sql.DB, user *appUser.AppUser) (bool, error) {
	sqlStmt := `
	INSERT INTO users (username, password)
	VALUES (?, ?);`
	_, err := db.Exec(sqlStmt, user.Username, hashPassword(user.Password))
	if err != nil {
		return false, err
	}
	return true, nil

}

func Validate(db *sql.DB, user *appUser.AppUser) (bool, error) {
	var hashedPassword string

	err := db.QueryRow("SELECT password FROM users WHERE username = ? ", user.Username).Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("utilisateur non trouvé")
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return false, fmt.Errorf("mot de passe incorrect")
	}

	return true, nil

}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)

}
