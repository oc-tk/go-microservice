package main

import (
	"log"
	"net/http"
	"os"

	"github.com/oc-tk/go_15.01/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	hg := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", hg)

	http.ListenAndServe(":9090", sm)
}
