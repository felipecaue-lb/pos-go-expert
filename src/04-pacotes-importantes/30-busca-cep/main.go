package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
		}

		f, err := os.Create("cidades.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer f.Close()

		_, err = f.WriteString(fmt.Sprintf("CEP: %s\nLogradouro: %s\nComplemento: %s\nUnidade: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nEstado: %s\nRegião: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n", data.Logradouro, data.Cep, data.Complemento, data.Unidade, data.Bairro, data.Localidade, data.Uf, data.Estado, data.Regiao, data.Ibge, data.Gia, data.Ddd, data.Siafi))
	}
}
