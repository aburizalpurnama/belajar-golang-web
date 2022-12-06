package test

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func formUpload(w http.ResponseWriter, r *http.Request) {
	err := MyTemplates.ExecuteTemplate(w, "upload.form", nil)
	if err != nil {
		panic(err)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan berupa file
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	written, err := io.Copy(fileDestination, file)
	fmt.Printf("written bytes copied: %v\n", written)
	if err != nil {
		panic(err)
	}

	// menerima inputan berupa value
	name := r.PostFormValue("name")
	MyTemplates.ExecuteTemplate(w, "upload.success", map[string]any{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", formUpload)
	mux.HandleFunc("/upload", uploadFile)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	s := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/xxx.jpg
var contohFile []byte

func TestUpload(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Rizal")

	file, err := writer.CreateFormFile("file", "gambar.png")
	assert.NoError(t, err)
	file.Write(contohFile)
	writer.Close()

	request, err := http.NewRequest(http.MethodPost, "http://localhost/", body)
	assert.NoError(t, err)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	uploadFile(recorder, request)

	bodyResponse, err := io.ReadAll(recorder.Result().Body)
	require.NoError(t, err)
	fmt.Println(string(bodyResponse))
}
