package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	error := t.Execute(os.Stdout, Cursos{{"Go", 40}, {"Python", 30}, {"JavaScript", 20}})
	if error != nil {
		panic(error)
	}
}
