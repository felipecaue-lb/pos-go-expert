package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var conta2 Conta
	err = json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		println(err)
	}
	println(conta2.Numero, conta2.Saldo)

	jsonPuro2 := []byte(`{"n":2,"s":200}`)
	var conta3 Conta
	err = json.Unmarshal(jsonPuro2, &conta3)
	if err != nil {
		println(err)
	}
	println(conta3.Numero, conta3.Saldo)
}
