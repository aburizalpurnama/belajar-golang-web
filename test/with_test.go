package test

import (
	"html/template"
	"net/http"
	"testing"
)

func withTemlate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	// t.ExecuteTemplate(w, "with.gohtml", map[string]any{
	// 	"Address": map[string]any{}},
	// )
	t.ExecuteTemplate(w, "with.gohtml", d)
}

func TestWithTemplate(t *testing.T) {
	runTest(t, withTemlate)
}
