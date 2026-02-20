package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	// Forma 1
	/* mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}) */

	// Forma 2
	/* mux.HandleFunc("/", HomeHandler) */

	// Forma 3
	mux.Handle("/", blog{title: "My Blog"})

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
