package main

import (
	"database/sql"
	"fmt"
	"github.com/xpdemon/userDb/database"
	"github.com/xpdemon/userDb/handle"
	"log"
	"net/http"
)

func main() {
	db := database.Initialize()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	http.HandleFunc("/addUser", func(w http.ResponseWriter, r *http.Request) {
		handle.AddUser(db, w, r)
	})

	http.HandleFunc("/validateUser", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validateUser reached")
		handle.ValidateUser(db, w, r)
	})

	err := http.ListenAndServe("0.0.0.0:8085", nil)
	if err != nil {
		panic(err)
	}

}
