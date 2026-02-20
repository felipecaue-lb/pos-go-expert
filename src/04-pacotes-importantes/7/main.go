package main

import (
	"fmt"
)

func main() {
	salarios := map[string]int{
		"Felipe":  1000,
		"Lucas":   2000,
		"Marcos":  3000,
		"Juliana": 4000}

	delete(salarios, "Lucas")

	salarios["Daiane"] = 5000
	salarios["Daiane 2"] = 15000

	for nome, salario := range salarios {
		fmt.Printf("Salário de %s é de %d\n", nome, salario)
	}

	salarios2 := make(map[string]int)
	salarios3 := map[string]int{}

	fmt.Println(salarios2)
	fmt.Println(salarios3)

	// Blank identifier
	for _, salario := range salarios {
		fmt.Printf("Salário de %d\n", salario)
	}
}
