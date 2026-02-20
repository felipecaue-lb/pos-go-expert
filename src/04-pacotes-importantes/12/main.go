package main

import (
	"fmt"
)

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	felipe := Cliente{
		Nome:  "Felipe",
		Idade: 30,
		Ativo: true,
	}

	felipe.Ativo = false

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", felipe.Nome, felipe.Idade, felipe.Ativo)
}
