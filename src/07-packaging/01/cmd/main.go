package main

import (
	"fmt"

	"github.com/felipecaue-lb/goexpert/07-packaging/1/01/math"
)

func main() {
	m := math.Math{A: 5, B: 3}
	fmt.Println(m.Add())
}
