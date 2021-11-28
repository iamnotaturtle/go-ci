package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Start() {
	r := CreateRouter()

	err := http.ListenAndServe("localhost:9000", middlewareTrimSlash(r))
	if err != nil {
		panic(err)
	}
}

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/character/{type}", homeHandler)
	http.Handle("/", r)
	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Character: %v\n", vars["type"])
}

func middlewareTrimSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
