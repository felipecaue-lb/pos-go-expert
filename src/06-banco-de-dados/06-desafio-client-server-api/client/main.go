package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type CurrencyRate struct {
	Bid float64 `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, error := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if error != nil {
		panic(error)
	}

	res, error := http.DefaultClient.Do(req)
	if error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout ao buscar cotação no server: contexto excedeu o limite de 300ms")
		}
		panic(error)
	}
	defer res.Body.Close()

	var currencyRate CurrencyRate
	error = json.NewDecoder(res.Body).Decode(&currencyRate)
	if error != nil {
		panic(error)
	}

	saveCurrencyRate(currencyRate.Bid)
}

func saveCurrencyRate(bid float64) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "Dólar: %.2f", bid)
	if err != nil {
		panic(err)
	}
}
