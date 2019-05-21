package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SomeOne visited our site")
	r.Header.Set("Content-Type", "text/html")
	fmt.Fprint(w, "<html><body>")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my sssuper awesome site!</h1>")
	} else if r.URL.Path =="/contact" {
		fmt.Fprint(w, "To get in touch, please send me email to <a href=\"mailto:o.dievri@outlook.com\">support</a><br>")

	}
	fmt.Fprint(w, r.URL.Path)
	fmt.Fprint(w, "</body></html>")
}