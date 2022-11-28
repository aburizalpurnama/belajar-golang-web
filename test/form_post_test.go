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

func formPost(w http.ResponseWriter, r *http.Request) {
	// Harus dilakukan parsing dulu
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hi, my name is %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	firstName := "Aburizal"
	lastName := "Purnama"

	requestBody := strings.NewReader(fmt.Sprintf("first_name=%s&last_name=%s", firstName, lastName))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	formPost(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	expected := fmt.Sprintf("Hi, my name is %s %s", firstName, lastName)
	require.Equal(t, expected, string(body))

	fmt.Printf("string(body): %v\n", string(body))
}
