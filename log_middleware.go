package belajargolangweb

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before execute handler")
	m.Handler.ServeHTTP(w, r)
	fmt.Println("After execute handler")
}
