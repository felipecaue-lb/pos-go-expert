package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado:", s)

	fmt.Println("A:", matematica.A)

	carro := matematica.Carro{Marca: "Fiat", Modelo: "Uno"}
	fmt.Println("Carro:", carro)

	fmt.Println(carro.Andar())
}
