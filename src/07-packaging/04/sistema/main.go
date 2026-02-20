/*
* Para baixar pacotes publicados e tento pacotes locais
* utilize: go mod tidy -e
 */

package main

import (
	"github.com/felipecaue-lb/goexpert/07-packaging/4/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
