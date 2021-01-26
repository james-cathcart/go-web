package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Instantiating server.")
	server := http.Server{
		Addr:    "127.0.0.1:443",
		Handler: nil,
	}

	log.Println("Starting server.")
	server.ListenAndServeTLS("cert.pem", "key.pem")

}
