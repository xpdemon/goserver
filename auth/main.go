// ext-authz/main.go

package main

import (
	"auth/handle"
	"fmt"
	"github.com/xpdemon/session"
	"github.com/xpdemon/session/cache"
	"net/http"
	"time"
)

func initApp() *cache.Cache {
	c := cache.NewCache()
	session.RenewKey(c)
	return c
}

func updateApp(c *cache.Cache) {
	session.RenewKey(c)
	fmt.Println("Key renewed")
}

func scheduleUpdateApp(cache *cache.Cache) {
	ticker := time.NewTicker(30 * time.Minute) // Intervalle de 30 minutes
	defer ticker.Stop()

	for range ticker.C {
		updateApp(cache)
		fmt.Println(cache)
	}
}

func main() {

	c := initApp()
	go scheduleUpdateApp(c)

	http.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticate reached")
		handle.Authenticate(w, r, c)
	})

	// Handler pour "/authorize" - utilisé par ext_authz pour l'authentification
	http.HandleFunc("/authorize/", func(w http.ResponseWriter, r *http.Request) {
		handle.Authorize(w, r, c)
	})

	fmt.Println("ext_authz service running on :9000")
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		fmt.Println("Erreur serveur:", err)
	}
}
