// Para ter todos os pacotes que a aplicação usa
// Utilize o comando: go mod vendor
// Dessa forma será criado uma pasta vendor com todos os pacotes utilizados

package main

import (
	"fmt"

	"github.com/felipecaue-lb/fc-utils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()

	fmt.Println(ed)
}
