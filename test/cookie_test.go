package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	// Create cookie
	name := r.URL.Query().Get("name")

	c := new(http.Cookie)
	c.Name = "X-my-name"
	c.Value = name
	c.Path = "/"

	if name != "" {
		http.SetCookie(w, c)
		fmt.Fprint(w, "Success set cookie")
	} else {
		fmt.Fprint(w, "name is not set")
	}
}

func TestSetCookie(t *testing.T) {
	name := "X-my-name"
	value := "Rizal"
	path := "/"
	targetURL := fmt.Sprintf("http://localhost:8081?name=%s", value)
	request := httptest.NewRequest(http.MethodPost, targetURL, nil)
	recorder := httptest.NewRecorder()

	setCookie(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.NotEmpty(t, string(body))

	cookies := response.Cookies()
	require.NotEmpty(t, cookies)

	fmt.Printf("body: %v\n", string(body))

	for i, c := range cookies {
		require.Equal(t, name, c.Name)
		require.Equal(t, value, c.Value)
		require.Equal(t, path, c.Path)
		fmt.Printf("%d: %v \n", i+1, c)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-my-name"
	cookie.Value = "Rizal"
	cookie.Path = "/"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	getCookie(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Println(string(body))
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("X-my-name")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "No Cookie")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome %s", c.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", setCookie)
	mux.HandleFunc("/get-cookie", getCookie)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
