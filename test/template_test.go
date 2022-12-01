package test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func simpleHtml(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	simpleHtml(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}

func simpleHtmlFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	simpleHtmlFile(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}

// pake pattern
func templateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	templateDirectory(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}

func templateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	templateEmbed(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}

/*
Mencoba memanggil template yang tidak ada
*/
func templateNoFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*gohtml"))
	t.ExecuteTemplate(w, "tidak-ada.gohtml", "Hello HTML Template")
}

func TestTemplateNoFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	templateNoFile(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}
