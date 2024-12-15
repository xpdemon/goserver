package main

import (
	"github.com/xpdemon/userDb/database"
	"github.com/xpdemon/userDb/handle"
	"net/http"
)

func main() {
	db := database.Initialize()

	http.HandleFunc("/addUser", func(w http.ResponseWriter, r *http.Request) {
		handle.AddUser(db, w, r)
	})

	err := http.ListenAndServe("0.0.0.0:8085", nil)
	if err != nil {
		panic(err)
	}

}
