package belajargolangweb

import (
	"fmt"
	"net/http"
)

type ErrorHandlerMiddleware struct {
	Handler http.Handler
}

func (m *ErrorHandlerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("RECOVER: ", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "ERROR: %s", err)
		}
	}()
	m.Handler.ServeHTTP(w, r)
}
