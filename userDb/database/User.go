package database

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Username string
	Password string
}

func New(username string, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

func Add(db *sql.DB, user *User) (bool, error) {
	sqlStmt := `
	INSERT INTO users (username, password)
	VALUES (?, ?);`
	_, err := db.Exec(sqlStmt, user.Username, hashPassword(user.Password))
	if err != nil {
		return false, err
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
