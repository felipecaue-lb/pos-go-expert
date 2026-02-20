package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://viacep.com.br/ws/38408383/json/")
	if err != nil {
		panic(err)
	}

	// Pula a execução dessa linha
	// Finalizar a execução do código completo
	// Executa essa linha
	defer req.Body.Close()

	println("Status Code:", req.StatusCode)

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println("Response Body:", string(res))
}
