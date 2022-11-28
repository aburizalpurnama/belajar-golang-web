package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Komponen yang berfungsi sebagai web server dan direpresentasikan dalam struct bernama Server
*/
func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8181",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
Komponen di dalam server yang menangani request yang masuk pada sebuah web server,
tanpa handler web server tidak bisa menangani request dari web browser/client.
*/
func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	}

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: handler,
	}

	server.ListenAndServe()
}

/*
- ServeMux berfungsi sebagai router atau suatu handler yang membungkus banyak endpoint di dalamnya.

  - Jika terdapat '/' pada akhir endpoint, maka semua endpoint turunannya kan mengacu pada endpoint pertama.
    akan tetapi, serverMux akan memproritaskan yang lebih panjang

    contoh : jika ada enpoint '/images/', maka jika kita mengakses '/images/thumbnail/' program akan mengarah ke
    function yang sama dengan '/images/' dengan catatan tidak ada enpoint '/images/thumbnail/',
    tapi jika ada maka akan mengakses function yang lebih panjang yaitu '/images/thumbnail/'.
*/
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi, my name is rizal")
	})
	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "this is IMAGES page")
	})
	mux.HandleFunc("/images/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "this is THUMBNAIL page")
	})

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	server.ListenAndServe()
}

/*
Kita dapat mengetahui isi dari request yang dikirim dari web client melalui function" yang ada pada argument http.Request
*/
func TestRequestExplore(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method:%s \nRequestURI:%s \nURL:%v \nHost:%s", r.Method, r.RequestURI, r.URL, r.Host)
	})

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := server.ListenAndServe()
	require.NoError(t, err)
}
