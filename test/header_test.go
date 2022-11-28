package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func requestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
	w.Header().Add("content-type", "plain-text")
}

func TestRequestHeader(t *testing.T) {
	contentType := "application/json"
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	request.Header.Add("content-type", contentType)
	recorder := httptest.NewRecorder()

	requestHeader(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	require.Equal(t, contentType, string(body))

	fmt.Printf("request content-type: %s \n", string(body))
	fmt.Printf("response content-type: %s \n", response.Header.Get("content-type"))
}
