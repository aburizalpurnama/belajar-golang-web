package test

import (
	"net/http"
	"testing"
)

/*
parsing template dilakukan di luar handler function, shg handler function hanya melakukan execute saja.
*/
func templateCahcing(w http.ResponseWriter, r *http.Request) {
	MyTemplates.ExecuteTemplate(w, "template-function", d)
}

func TestTemplateCaching(t *testing.T) {
	runTest(t, templateCahcing)
}
