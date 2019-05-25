package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.ExecuteTemplate(w, "home.gohtml", nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.ExecuteTemplate(w, "contact.gohtml", nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml", "views/contact.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", r)
}