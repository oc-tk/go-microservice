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

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")
	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Error occurred", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "hello %s", d)
}
