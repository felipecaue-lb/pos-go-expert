package main

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T Number](a, b T) bool {
	return a == b
}

/* func Compara[T comparable](a, b T) bool {
	return a == b
} */

func main() {
	m := map[string]int{"Felipe": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m))

	m2 := map[string]float64{"Felipe": 100.10, "João": 200.20, "Maria": 300.30}
	println(Soma(m2))

	m3 := map[string]MyNumber{"Felipe": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m3))

	println(Compara(1, 1.1))
}
