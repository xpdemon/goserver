package handle

import (
	"database/sql"
	"fmt"
	"github.com/xpdemon/userDb/appUser"
	"github.com/xpdemon/userDb/database"
	"io"
	"net/http"
	"strings"
)

func AddUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	userString := readBody(r)

	u := appUser.New(strings.Split(userString, ":")[0], strings.Split(userString, ":")[1])

	_, err := database.Add(db, u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("addUser : %v\n", u)

	w.WriteHeader(http.StatusOK)
}

func ValidateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	userString := readBody(r)
	u := appUser.New(strings.Split(userString, ":")[0], strings.Split(userString, ":")[1])

	_, err := database.Validate(db, u)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Printf("user : %v is invalid\n", u.Username)
		return
	}

	fmt.Printf("user : %v is valid\n", u.Username)

}

func readBody(r *http.Request) string {
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

	return string(body)
}
