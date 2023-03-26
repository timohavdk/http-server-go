package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	server.ListenAndServe()
}
