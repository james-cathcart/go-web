package main

import (
	"html/Template"
	"net/http"
)

func main() {

	mux := http.NewServeMux() // create multiplexer

	files := http.FileServer(http.Dir("/public")) // serve static files from a particular directory

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index) // ok book, wheres the index declaration?
	mux.HandleFunc("/err", err)
	mux.HandleFunc("login", login)
	mux.HandleFunc("login", login)
	mux.HandleFunc("login", login)
	mux.HandleFunc("login", login)
	mux.HandleFunc("login", login)
	mux.HandleFunc("login", login)
	mux.HandleFunc("login", login)
	mux.HandleFunc("login", login)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()

}

func index(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	threads, err := data.threads()

	if err != nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

