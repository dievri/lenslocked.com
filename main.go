package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w,"To get in touch, please send me email to <a href=\"mailto:o.dievri@outlook.com\">support</a>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprint(w, "The page you looking for is not found")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "NEVER GIVE UP!")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", r)
}