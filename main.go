package main

import (
	"fmt"
	"net/http"
)

type Index struct {
}

func (h *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index page")
}

type About struct {
}

func (h *About) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page")
}

func main() {
	index := Index{}
	about := About{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/", &index)
	http.Handle("/about/", &about)
	server.ListenAndServe()
}
