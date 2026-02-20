package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(40, 20)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior ou igual a 50")
	}

	return a + b, nil
}
