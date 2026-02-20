package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// CRIAÇÃO DE ARQUIVOS
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// ESCRITA DE ARQUIVOS
	tamanho, err := f.Write([]byte("Escrevendo no arquivo"))
	//tamanho, err := f.WriteString("Escrevendo no arquivo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tamanho do arquivo: %d bytes\n", tamanho)

	f.Close()

	// LEITURA DE ARQUIVOS
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Conteúdo do arquivo:", string(arquivo))

	// LEITURA DE ARQUIVOS COM BUFFERS
	f, err = os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	buffer := make([]byte, 10) // Lê 10 bytes por vez
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break // Fim do arquivo
			}
			panic(err)
		}
		fmt.Println(string(buffer[:n])) // Imprime os bytes lidos
	}

	// REMOVER UM ARQUIVO
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo removido com sucesso.")
}
