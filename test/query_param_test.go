package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello, welcome to golang web")
	} else {
		fmt.Fprintf(w, "Hello %s, welcome to golang web", name)
	}
}

func TestSayHello(t *testing.T) {
	name := "Rizal"
	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:8081?name=%s", name), nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	require.Equal(t, fmt.Sprintf("Hello %s, welcome to golang web", name), string(body))

	fmt.Printf("string(body): %v\n", string(body))
}

func multipleParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	fmt.Fprintf(w, "Hi, my name is %s %s", firstName, lastName)
}

func TestMultipleParameter(t *testing.T) {
	firstName := "Rizal"
	lastName := "Purnama"
	targetURL := fmt.Sprintf("http://localhost:8081?first_name=%s&last_name=%s", firstName, lastName)

	request := httptest.NewRequest(http.MethodPost, targetURL, nil)
	recorder := httptest.NewRecorder()

	multipleParameter(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	require.Equal(t, fmt.Sprintf("Hi, my name is %s %s", firstName, lastName), string(body))
	fmt.Printf("string(body): %v\n", string(body))
}

func multipleValueParameter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]

	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleValueParameter(t *testing.T) {
	names := []string{"Moh.", "Aburizal", "Purnama"}
	targetURL := fmt.Sprintf("http://localhost:8081?name=%s&name=%s&name=%s", names[0], names[1], names[2])
	request := httptest.NewRequest(http.MethodGet, targetURL, nil)
	recorder := httptest.NewRecorder() 

	multipleValueParameter(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	expected := fmt.Sprint(strings.Join(names, " "))
	require.Equal(t, expected, string(body))

	fmt.Printf("string(body): %v\n", string(body))
}
