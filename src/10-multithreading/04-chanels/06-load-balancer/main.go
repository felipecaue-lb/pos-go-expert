package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtdWorkers := 10

	for i := range qtdWorkers {
		go worker(i, data)
	}

	for i := range 100 {
		data <- i
	}
}
