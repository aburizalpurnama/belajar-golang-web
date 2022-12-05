package test

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type data struct {
	Title   string
	Name    string
	Address address
	Score   int
	Hobbies []string
}

func (d data) SayHello(name string) string {
	return fmt.Sprintf("Hello %s, my name is %s and my score %d", name, d.Name, d.Score)
}

type address struct {
	City   string
	Street string
}

/*
Using golang embed
*/
//go:embed templates/*.gohtml
var templates embed.FS

var MyTemplates = template.Must(template.ParseFS(templates, "templates/*gohtml"))

var d data = data{
	Title: "Belajar Golang Web",
	Name:  "Rizal",
	Address: address{
		City:   "Probolinggo",
		Street: "Jl. Cempaka"},
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func runTest(t *testing.T, f func(w http.ResponseWriter, r *http.Request)) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	f(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	fmt.Printf("string(body): %v\n", string(body))
}
