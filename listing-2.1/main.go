package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux() // create multiplexer

	files := http.FileServer(http.Dir("/public")) // serve static files from a particular directory

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index) // ok book, wheres the index declaration?

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()

}
