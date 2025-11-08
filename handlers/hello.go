package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")

	data, error := io.ReadAll(r.Body)

	if error != nil {
		http.Error(rw, "o", http.StatusBadRequest)

		return
	}

	fmt.Fprintf(rw, "Hello %s\n", data)
	fmt.Println(r)
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
