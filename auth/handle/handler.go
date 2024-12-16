package handle

import (
	"fmt"
	"github.com/xpdemon/session"
	"github.com/xpdemon/session/cache"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	// Générer un session id + signature
	id, er := session.GenerateSessionID(32)
	if er != nil {
		fmt.Println("Erreur génération ID:", er)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	signed := sign(id, c)

	// Générer un cookie de session signé
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: signed,
		Path:  "/",
	})
	return

}

func sign(id string, c *cache.Cache) string {

	sid := session.SignID(id, c)
	fmt.Println("id: " + id)
	fmt.Println("signed id: " + sid)

	return sid
}

func Authorize(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
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
