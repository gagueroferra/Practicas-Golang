package main

import "testing"

func TestImpuestos(t *testing.T) {
	salario := float32(510000)

	resultado_esperado := 86700

	if Impuestos(salario) != resultado_esperado {
		t.Fatal("No es el resultado esperado")
	}
}
