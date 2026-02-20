package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))

		error := t.Execute(w, Cursos{
			{"Go", 40},
			{"Python", 30},
			{"JavaScript", 20},
		})
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
