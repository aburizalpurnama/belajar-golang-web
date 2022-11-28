package test

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

/*
Langsung masuk ke folder yang diembed, shg dapat langsung mengakses sbb:
http://localhost:8081/static/index.html
*/
func TestFileServerWithEmbedSubDir(t *testing.T) {
	dir, err := fs.Sub(resources, "resources")
	assert.NoError(t, err)

	var fileserver = http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
Foldernya juga ikut diembed, sehingga harus menyebutkan nama foldernya sbb:
http://localhost:8081/static/resources/index.html
*/
func TestFileServerWithEmbed(t *testing.T) {
	var fileserver = http.FileServer(http.FS(resources))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
Mengakses folder, tapi tidak dibundle ke execute file. jadi harus dicopy secara manual ke server
*/
func TestFileServer(t *testing.T) {
	var dir = http.Dir("./../resources")
	var fileserver = http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
Dapat secara dinamis mengembalikan file
*/

func serveFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(serveFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func serveFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {

	s := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(serveFileEmbed),
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
