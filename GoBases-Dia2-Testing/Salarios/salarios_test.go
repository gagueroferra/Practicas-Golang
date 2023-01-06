package main

import "testing"

func TestSalario(t *testing.T) {
	tiempo := 10000
	cat := "A"
	resultado_esperado := 747000

	if Salario(tiempo, cat) != float32(resultado_esperado) {
		t.Fatal("El resultado no es el esperado")
	}
}
