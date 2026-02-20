package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
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
