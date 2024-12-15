package handle

import (
	"database/sql"
	"fmt"
	"github.com/xpdemon/userDb/database"
	"io"
	"net/http"
	"strings"
)

func AddUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Erreur fermeture body:", err)
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Erreur lecture body:", err)
	}

	userString := string(body)
	user := database.User{
		Username: strings.Split(userString, ":")[0],
		Password: strings.Split(userString, ":")[1],
	}

	_, err = database.Add(db, &user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("addUser : %v\n", user)

	w.WriteHeader(http.StatusOK)
}
