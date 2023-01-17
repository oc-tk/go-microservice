package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

type Goodbye struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Error occurred", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "hello %s", d)
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee"))
	g.l.Printf("Byee")
}
