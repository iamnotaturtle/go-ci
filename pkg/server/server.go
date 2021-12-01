package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Starts an http server on :9000
func Start() {
	r := CreateRouter()

	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "localhost:9000"
	} else {
		port = ":" + port
	}

	fmt.Printf("Listening on port %s\n", port)

	err := http.ListenAndServe(port, middlewareTrimSlash(r))
	if err != nil {
		panic(err)
	}
}

// Creates a router and adds routes
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/character/{type}", homeHandler)
	return r
}

// Handles root requests
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "it works I promise!")
}

// Handles incoming requests
func homeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Character: %v\n", vars["type"])
}

// Trims any trailing slashes on urls.
// Otherwise we would 404 on a trailing slash.
func middlewareTrimSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}
