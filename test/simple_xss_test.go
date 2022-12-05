package test

import (
	"html/template"
	"net/http"
	"testing"
)

/*
secara default, html template sudah aman dari XSS.
*/
func simpleXSS(w http.ResponseWriter, r *http.Request) {
	MyTemplates.ExecuteTemplate(w, "sample-xss", map[string]any{
		"Title": "Contoh Cross Site Scripting",
		"Body":  "ini adalah body content <h1> You have been hacked</h1>",
	})
}

func simpleXSSDisabled(w http.ResponseWriter, r *http.Request) {
	MyTemplates.ExecuteTemplate(w, "sample-xss", map[string]any{
		"Title": "Contoh Cross Site Scripting",
		"Body":  template.HTML("ini adalah body content <h1> You have been hacked</h1>"),
	})
}

func TestSimpleXSS(t *testing.T) {
	runTest(t, simpleXSS)
}

func TestSimpleXSSServer(t *testing.T) {
	s := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(simpleXSS),
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSimpleXSSDisabledServer(t *testing.T) {
	s := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(simpleXSSDisabled),
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func simpleXSSDisabledParam(w http.ResponseWriter, r *http.Request) {
	MyTemplates.ExecuteTemplate(w, "sample-xss", map[string]any{
		"Title": "Contoh Cross Site Scripting",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

/*
akses url berikut pada browser:
http://localhost:8081/?body=<h1>Welcome</h1><script>alert('You have been hacked')</script>

maka aplikasi akan merender code tsb. (bisa dilakukan cross site scripting)
*/
func TestSimpleXSSDisabledParamServer(t *testing.T) {
	s := http.Server{
		Addr:    "localhost:8081",
		Handler: http.HandlerFunc(simpleXSSDisabledParam),
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
