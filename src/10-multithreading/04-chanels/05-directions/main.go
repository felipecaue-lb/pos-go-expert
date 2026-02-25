package main

import "fmt"

func recebe(nome string, chanel chan<- string) {
	chanel <- nome
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	chanel := make(chan string)
	go recebe("Felipe", chanel)
	ler(chanel)
}
