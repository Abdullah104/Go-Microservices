package handlers

import (
	"net/http"
)

type Goodbye struct{}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Good bye"))
}

func NewGoodbye() *Goodbye {
	return &Goodbye{}
}
