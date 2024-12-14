// ext-authz/main.go

package main

import (
	"fmt"
	"github.com/xpdemon/session"
	"net/http"
)

func main() {
	// Handler pour "/authz" - utilisé par ext_authz pour l'authentification
	http.HandleFunc("/authz/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("you are on /authz")

		cookie, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println("Erreur cookie:", err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		_, e := session.ValidateSignedID(cookie.Value)
		if e != nil {
			fmt.Println("Erreur session:", e)
			// Session invalide ou falsifiée
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Session valide
		w.WriteHeader(http.StatusOK)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	fmt.Println("ext_authz service running on :9000")
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		fmt.Println("Erreur serveur:", err)
	}
}
