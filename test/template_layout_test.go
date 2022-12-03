package test

import (
	"html/template"
	"net/http"
	"testing"
)

func templateLayout(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "simple-content", d)
}

func TestTemplateLayout(t *testing.T) {
	runTest(t, templateLayout)
}
