/*
* Para testar, use o comando `go test`
* Para ver mais detalhes no teste, use `go test -v`
* Para testar a cobertura de teste, use `go test -coverprofile=coverage.out` e depois `go tool cover -html=coverage.out`
* Se estiver trabalhando dentro do container, use `go tool cover -html=coverage.out -o coverage.html` para gerar o html
* Para rodar o benchmark, use `go test -bench=.`
* Para rodar somente o benchmark, use `go test -bench=. -run=^#`
* Para rodar o fuzz test, use `go test -fuzz=. -run=^#`
 */

package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("CalculateTax(%f) = %f; expected %f", amount, result, expected)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{amount: 500.0, expect: 5.0},
		{amount: 1000.0, expect: 10.0},
		{amount: 1500.0, expect: 10.0},
		{amount: 0.0, expect: 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expect {
			t.Errorf("CalculateTax(%f) = %f; expected %f", item.amount, result, item.expect)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 0.0, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0.0 {
			t.Errorf("Reseived %f but expected 0.0", result)
		}
		if amount > 20000 && result != 20.0 {
			t.Errorf("Reseived %f but expected 20.0", result)
		}
	})
}
