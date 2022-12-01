package test

import (
	"html/template"
	"net/http"
	"testing"
)

func templateComparator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	d.Score = 50
	t.ExecuteTemplate(w, "ge-comparable.gohtml", d)
}

func TestTemplateEqComparator(t *testing.T) {
	runTest(t, templateComparator)
}
