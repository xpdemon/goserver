package handle

import (
	"fmt"
	"github.com/xpdemon/session"
	"github.com/xpdemon/session/cache"
	"io"
	"net/http"
)

func Sign(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
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
	id := string(body)
	sid := session.SignID(id, c)
	fmt.Println("id: " + id)
	fmt.Println("signed id: " + sid)
	_, e := w.Write([]byte(sid))
	if e != nil {
		fmt.Println(e.Error())
	}
}

func Auth(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println("Erreur cookie:", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	_, e := session.ValidateSignedID(cookie.Value, c)
	if e != nil {
		fmt.Println("Erreur session:", e)
		// Session invalide ou falsifiée
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Session valide
	w.WriteHeader(http.StatusOK)

}
