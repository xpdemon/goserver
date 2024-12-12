package main

import (
	"auth/sessionId"
	"fmt"
	"net/http"
)

// Clé secrète utilisée pour signer les sessions.
// Dans un vrai déploiement, stockez cette clé en variable d'environnement ou secret.

func main() {
	// Handler pour "/auth" qui redirige vers "/auth/"
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		// Redirige de "/auth" vers "/auth/" en 301 (Moved Permanently)
		fmt.Println(r.RequestURI + " Auth 01 \n")
		_, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println("Erreur cookie:", err)
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
	})

	// Handler principal pour "/auth/" - gère /auth/ et tout ce qui suit
	http.HandleFunc("/auth/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI + "\n")

		cookie, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println("Erreur cookie:", err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		sessionID, e := sessionId.ValidateSignedSessionID(cookie.Value)
		if e != nil {
			fmt.Println("Erreur session:", e)
			// Session invalide ou falsifiée
			http.Redirect(w, r, "/login", http.StatusForbidden)
			return
		}

		// Session valide
		fmt.Fprintf(w, "Session valide: %s", sessionID)
	})

	// Handler pour "/login"
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Afficher un formulaire de login
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
			}
			user := r.Form.Get("username")
			pass := r.Form.Get("password")

			if user == "admin" && pass == "secret" {
				// Générer un session id + signature
				sid, er := sessionId.GenerateSessionID(32)
				if er != nil {
					fmt.Println("Erreur génération ID:", er)
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}

				signed := sessionId.SignSessionID(sid)

				// Générer un cookie de session signé
				http.SetCookie(w, &http.Cookie{
					Name:  "session_id",
					Value: signed,
					Path:  "/",
				})
				http.Redirect(w, r, "/auth", http.StatusFound)
			} else {
				http.Redirect(w, r, "/login", http.StatusForbidden)
			}
		}
	})

	fmt.Println("ext_authz service running on :9000")
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		fmt.Println("Erreur serveur:", err)
	}
}
