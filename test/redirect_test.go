package test

import (
	"fmt"
	"net/http"
	"testing"
)

func destinationRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You've been redirected")
}

func originHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect", http.StatusTemporaryRedirect)
}

func TestRedirections(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/origin", originHandler)
	mux.HandleFunc("/redirect", destinationRedirect)

	s := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
