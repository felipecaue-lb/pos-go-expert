package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	var cep string = "38408383"

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		//time.Sleep(time.Second * 4)

		resp1, error := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
		if error != nil {
			panic(error)
		}
		defer resp1.Body.Close()

		body, error := io.ReadAll(resp1.Body)
		if error != nil {
			panic(error)
		}

		c1 <- string(body)
	}()

	go func() {
		//time.Sleep(time.Second * 4)

		resp2, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if error != nil {
			panic(error)
		}
		defer resp2.Body.Close()

		body, error := io.ReadAll(resp2.Body)
		if error != nil {
			panic(error)
		}

		c2 <- string(body)
	}()

	select {
	case msg := <-c1:
		fmt.Printf("BrasilAPI => %s\n", msg)

	case msg := <-c2:
		fmt.Printf("ViaCEP => %s\n", msg)

	case <-time.After(time.Second):
		println("Timeout: Nenhuma resposta recebida dentro de 3 segundos")
	}

}
