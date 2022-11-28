package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Agar dapat dilakukan unit test, sebaiknya jangan benar" menjalankan server. tetapi menggunakan package httptest.

httptest.NewRequest() untuk membuat object request
httptest.NewRecorder() untuk membuat object yang merekan response dari web server
httptest.NewRecorder().Result() untuk mendapatkan response dari web server
*/

var bodyVal = "Hello World!"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, bodyVal)
}

func TestHelloHandler(t *testing.T) {
	method := http.MethodGet
	URL := "http://localhost:8081/hello"
	request := httptest.NewRequest(method, URL, nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	require.NoError(t, err)
	require.Equal(t, bodyVal, string(body))

	fmt.Printf("response.StatusCode: %v\n", response.StatusCode)
	fmt.Printf("response.Status: %v\n", response.Status)
	fmt.Printf("string(body): %v\n", string(body))
}
