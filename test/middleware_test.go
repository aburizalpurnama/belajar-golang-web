package test

import (
	belajargolangweb "belajar-golang-web"
	"fmt"
	"net/http"
	"testing"
)

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you have been access middleware")
	})
	mux.HandleFunc("/others", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, it's another one")
	})

	LogMiddleware := new(belajargolangweb.LogMiddleware)
	LogMiddleware.Handler = mux

	s := http.Server{
		Addr:    "localhost:8081",
		Handler: LogMiddleware,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
