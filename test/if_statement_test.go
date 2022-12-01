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

func ifStatement(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", data{
		Title: "Template Action (If-Statement)",
		Name:  "Hasbi",
	})
}

func TestIfStatement(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	ifStatement(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}
