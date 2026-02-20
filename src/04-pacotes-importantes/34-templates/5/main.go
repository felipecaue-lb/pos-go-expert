package main

import (
	"net/http"
	"text/template"
)

func main() {
	templates := []string{"header.html", "content.html", "footer.html"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		error := t.Execute(w, nil)
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
