package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)

		waitGroup.Done() // Marca a tarefa como concluída
	}
}

// Thread 1 - main já é uma thread
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25) // Adiciona 25 tarefas ao WaitGroup

	// Thread 2
	go task("A", &waitGroup)

	// Thread 3
	go task("B", &waitGroup)

	// Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "Anonymous")
			time.Sleep(1 * time.Second)

			waitGroup.Done() // Marca a tarefa anônima como concluída
		}
	}()

	waitGroup.Wait() // Espera todas as tarefas terminarem
}
