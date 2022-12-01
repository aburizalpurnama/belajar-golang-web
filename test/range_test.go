package test

import (
	"html/template"
	"net/http"
	"testing"
)

func templateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	d.Hobbies = []string{"Music", "Gadget", "Movie"}
	t.ExecuteTemplate(w, "range.gohtml", d)
}

func TestTemplateRange(t *testing.T) {
	runTest(t, templateRange)
}
