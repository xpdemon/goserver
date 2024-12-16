package main

import (
	"fmt"
	"github.com/xpdemon/session"
	"io"
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

			response, err := http.Post("http://127.0.0.1:8085", "", strings.NewReader(userPwd))
			if err != nil {
				fmt.Println("impossible de joindre la base de données")
			}

			switch response.StatusCode {
			case http.StatusOK:
				generateSid(w, r)
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

func generateSid(w http.ResponseWriter, r *http.Request) {
	// Générer un session id + signature
	sid, er := session.GenerateSessionID(32)
	if er != nil {
		fmt.Println("Erreur génération ID:", er)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	fmt.Println("sid: " + sid)
	signed := signId(sid)
	fmt.Println("signed: " + signed)

	// Générer un cookie de session signé
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: signed,
		Path:  "/",
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func signId(id string) string {
	r := strings.NewReader(id)
	fmt.Println(" send request to authz")
	resp, err := http.Post("http://0.0.0.0:9000/sign", "", r)
	if err != nil {
		return fmt.Sprint(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Erreur fermeture body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprint(err)
	}

	fmt.Println(" got response from authz" + string(body))

	return string(body)

}
