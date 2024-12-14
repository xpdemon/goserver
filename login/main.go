package main

import (
	"fmt"
	"github.com/xpdemon/session"
	"net/http"
)

func main() {
	// Handler principal pour "/" - protégé
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("you are on /")

		_, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println("Erreur cookie:", err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// La validation est gérée par Envoy via ext_authz
		// Ici, vous pouvez ajouter des fonctionnalités supplémentaires si nécessaire
		fmt.Fprintln(w, "Bienvenue, vous êtes authentifié !")
	})

	// Handler pour "/login"
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		if r.Method == http.MethodGet {
			fmt.Println("print form")
			w.Write([]byte(`
                <html>
                <body>
                  <form method="POST">
                    <input type="text" name="username" placeholder="Username">
                    <input type="password" name="password" placeholder="Password">
                    <input type="submit" value="Login">
                  </form>
                </body>
                </html>
            `))
			return
		}

		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				fmt.Println("Erreur de parsing:", err)
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			user := r.Form.Get("username")
			pass := r.Form.Get("password")

			if user == "admin" && pass == "secret" {
				// Générer un session id + signature
				sid, er := session.GenerateSessionID(32)
				if er != nil {
					fmt.Println("Erreur génération ID:", er)
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}

				signed := session.SignID(sid)

				// Générer un cookie de session signé
				http.SetCookie(w, &http.Cookie{
					Name:  "session_id",
					Value: signed,
					Path:  "/",
				})
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				// Afficher un message d'erreur ou rediriger
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		}
	})

	fmt.Println("go_app service running on :8081")
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		fmt.Println("Erreur serveur:", err)
	}
}
