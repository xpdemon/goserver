package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

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

			userPwd := strings.Join([]string{r.Form.Get("username"), r.Form.Get("password")}, ":")

			response, err := http.Post("http://0.0.0.0:8085/validateUser", "", strings.NewReader(userPwd))
			if err != nil {
				fmt.Println("impossible de joindre la base de données")
			}

			switch response.StatusCode {
			case http.StatusOK:
				resp, err := http.Get("http://0.0.0.0:9000/authenticate")
				if err != nil {
					fmt.Println(err)
				}

				for _, cook := range resp.Cookies() {
					http.SetCookie(w, cook)

				}
				http.Redirect(w, r, "/", http.StatusFound)

			case http.StatusForbidden:
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
