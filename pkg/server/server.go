package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const port = 9000

// Starts an http server on :9000
func Start() {
	r := CreateRouter()

	http.Handle("/", r)

	fmt.Printf("Listening on port:%d\n", port)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), middlewareTrimSlash(r))
	if err != nil {
		panic(err)
	}
}

// Creates a router and adds routes
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/character/{type}", homeHandler)
	return r
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
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
