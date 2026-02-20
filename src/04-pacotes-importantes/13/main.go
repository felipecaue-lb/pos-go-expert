package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	felipe := Cliente{
		Nome:  "Felipe",
		Idade: 30,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua das Flores",
			Numero:     456,
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}

	felipe.Ativo = false
	felipe.Cidade = "Bady Bassitt"

	fmt.Printf("Nome: %s\nIdade: %d\nEndereço: %s, %d - %s/%s\nAtivo: %t\n", felipe.Nome, felipe.Idade, felipe.Logradouro, felipe.Numero, felipe.Cidade, felipe.Estado, felipe.Ativo)
}
