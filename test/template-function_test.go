package test

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"testing"
)

func templateFunction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	d.Score = 89
	t.ExecuteTemplate(w, "template-function", d)
}

func TestTemplateFunction(t *testing.T) {
	runTest(t, templateFunction)
}

func templateGlobalFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("template-function").Parse(`{{ len .Name}}`))
	t.ExecuteTemplate(w, "template-function", d)
}

func TestTemplateGlobalFunction(t *testing.T) {
	runTest(t, templateGlobalFunction)
}

func templateCreateGlobalFunction(w http.ResponseWriter, r *http.Request) {
	t := template.New("template-function")
	t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t.Parse(`{{ upper .Name}}`)
	t.ExecuteTemplate(w, "template-function", d)
}

func TestTemplateCreateGlobalFunction(t *testing.T) {
	runTest(t, templateCreateGlobalFunction)
}

func templateFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("template-function")
	t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
		"sayHello": func(name1, name2 string) string {
			return fmt.Sprintf("Hello %s, my name is %s", name1, name2)
		},
	})

	t.Parse(`{{ upper .Name | sayHello "Zakir"}}`)
	t.ExecuteTemplate(w, "template-function", d)
}

func TestTemplateFunctionPipeline(t *testing.T) {
	runTest(t, templateFunctionPipeline)
}
