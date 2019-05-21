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
	fmt.Fprint(w, "<h1>Welcome to my sssuper awesome site!</h1>")
}