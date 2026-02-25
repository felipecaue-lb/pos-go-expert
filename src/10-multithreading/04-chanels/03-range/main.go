package main

import "fmt"

// Thread 1
func main() {
	chanel := make(chan int)

	go publish(chanel)
	reader(chanel)
}

func reader(chanel chan int) {
	for i := range chanel {
		fmt.Printf("Received %d\n", i)
	}
}

func publish(chanel chan int) {
	for i := 0; i < 10; i++ {
		chanel <- i
	}

	close(chanel)
}
