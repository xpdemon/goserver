package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/auth/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request receive %v\n", r.RequestURI)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader, "Basic ") {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		encoded := strings.TrimPrefix(authHeader, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		user, pass := parts[0], parts[1]
		if user == "admin" && pass == "secret" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	})

	fmt.Println("ext_authz service running on :9000")
	http.ListenAndServe("0.0.0.0:9000", nil)
}
