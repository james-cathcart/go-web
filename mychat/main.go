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

func index(writer http.ResponseWriter, request *http.Request) {

	threads, err := data.Threads()
	if err == nil {
		_, err := session(writer, request)
	}

	public_tmpl_files := []string{"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html"}
	private_tmpl_files := []string{"templates/layout.html",
		"templates/private.navbar.html",
		"templates/index.html"}

	var templates *template.Template

	if err != nil {
		templates = template.Must(template.ParseFiles(private_tmpl_files...))
	} else {
		templates = template.Must(template.ParseFiles(public_tmpl_files...))
	}

	templates.ExecuteTemplate(writer, "layout", threads)
}

func authenticate(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, _ := data.UserByEmail(r.PostFormValue("email"))

	if user.Password == data.Encrypt(r.PostFormValue("password")) {

		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)

	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
