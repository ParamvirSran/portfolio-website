package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handle root url
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// handle msg url for htmx requests
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>Hello, this is my portfolio website to showcase my software engineering journey!</p>")
	})

	// serve static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start server
	fmt.Println("Server is running at http://localhost:80")
	http.ListenAndServe(":80", nil)
}
