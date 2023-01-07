package main

import "testing"

func TestCalculadora(t *testing.T) {
	operacion := Promedia

	resultado_esperado := float32(5.125)

	calculo, err := Estadistica(operacion)

	operation := calculo(10, 9, 5, 2, 5, 1, 6, 3)

	if operation != resultado_esperado {
		t.Fatal("El resultado no es el esperado")
	} else if err != nil {
		panic(err)
	}
}
