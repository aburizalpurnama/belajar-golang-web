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

type data struct {
	Title   string
	Name    string
	Address address
}

type address struct {
	City   string
	Street string
}

func templateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name-value.gohtml"))
	t.ExecuteTemplate(w, "name-value.gohtml",
		data{
			Title: "Learning Go data struct template",
			Name:  "Rizal"})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	templateDataStruct(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}

func templateDataNestedStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/nested-struct.gohtml"))
	t.ExecuteTemplate(w, "nested-struct.gohtml",
		data{
			Title: "Learning Go data struct template",
			Name:  "Rizal",
			Address: address{
				City:   "Jakarta Selatan",
				Street: "Jl. Wijaya kusuma",
			}})
}

func TestTemplateDataNestedStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	templateDataNestedStruct(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}

func templateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name-value.gohtml"))
	t.ExecuteTemplate(w, "name-value.gohtml", map[string]any{
		"Title": "Learning Go data map template",
		"Name":  "Moh. Aburizal Purnama",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	templateDataMap(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}
