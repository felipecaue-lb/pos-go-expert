package main

// Thread 1
func main() {
	canal := make(chan string) // Canal criado, sem buffer

	// Thread 2
	go func() {
		canal <- "Olá, mundo!" // Bloqueia até alguém ler
	}()

	// Thread 1
	mensagem := <-canal // Bloqueia até alguém enviar
	println(mensagem)
}
