package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	chanel := make(chan int)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)

	go publish(chanel)
	reader(chanel, &waitGroup)

	waitGroup.Wait()
}

func reader(chanel chan int, waitGroup *sync.WaitGroup) {
	for i := range chanel {
		fmt.Printf("Received %d\n", i)
		waitGroup.Done()
	}
}

func publish(chanel chan int) {
	for i := 0; i < 10; i++ {
		chanel <- i
	}

	close(chanel)
}
