package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if err != nil {
		fmt.Printf("An Error occurred : %v", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	user := decodeHeader(authHeader)

	fmt.Printf("got / request \n")

	s := strings.Join([]string{"This is the envoy proxy poc", user}, ":")
	_, err := io.WriteString(w, s)
	if err != nil {
		fmt.Println("Oops")
	}

}

func getHello(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	user := decodeHeader(authHeader)

	fmt.Printf("got / request \n")

	s := strings.Join([]string{"Hello ", user}, ":")
	_, err := io.WriteString(w, s)
	if err != nil {
		fmt.Println("Oops")
	}

}

func decodeHeader(authHeader string) string {
	encoded := strings.TrimPrefix(authHeader, "Basic ")
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("oops")
	}
	parts := strings.SplitN(string(decoded), ":", 2)
	user := parts[0]

	return user
}
