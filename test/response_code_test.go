package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func responseStatusCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is empty")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hi, %s", name)
	}
}

func TestResponseStatusCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	responseStatusCode(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	statusCode := response.StatusCode
	status := response.Status

	fmt.Printf("statusCode: %v\n", statusCode)
	fmt.Printf("status: %v\n", status)
	fmt.Printf("body: %v\n", string(body))
}
