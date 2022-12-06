package test

import (
	"fmt"
	"net/http"
	"testing"
)

func downloadRendered(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Tidak ada nama file")
		return
	}

	w.Header().Add("Content-Disposition", "inline")
	http.ServeFile(w, r, "./resources/"+fileName)
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Tidak ada nama file")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(w, r, "./resources/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/download", downloadFile)
	mux.HandleFunc("/render", downloadRendered)

	s := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
