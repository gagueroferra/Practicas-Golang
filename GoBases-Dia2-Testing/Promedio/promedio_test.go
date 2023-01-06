package main

import "testing"

func TestPromedio(t *testing.T) {
	lista_notas := []int{2, 4, 6, 8, 10}

	resultado_esperado := 6

	if Promedio(lista_notas...) != float32(resultado_esperado) {
		t.Fatal("El resultado no es el esperado")
	}
}
