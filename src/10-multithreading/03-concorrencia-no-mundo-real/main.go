/* Para verificar problemas de race condition, podemos executar a aplicação com `go run -race main.go`
 * e acessar a página várias vezes. O Go irá detectar se há condições de corrida e exibir mensagens de erro indicando onde elas ocorrem.
 * Pode também utilizar a ferramenta apache bench para simular múltiplas requisições simultâneas e verificar o comportamento da aplicação.
 * Exemplo de comando para simular 100 requisições com 10 concorrentes:
 * `ab -n 100 -c 10 http://localhost:8080/`
 * Isso irá enviar 100 requisições para a aplicação com um nível de concorrência de 10, permitindo observar como a aplicação lida com múltiplas requisições simultâneas.
 */

package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	//mutex := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1° Opção: Com mutex
		//mutex.Lock()
		//number++
		//mutex.Unlock()

		// 2° Opção: Com atomic
		atomic.AddUint64(&number, 1)

		time.Sleep(300 * time.Millisecond)
		fmt.Fprintf(w, "Você teve acesso a essa página %d vezes\n", number)
	})

	http.ListenAndServe(":8080", nil)
}
