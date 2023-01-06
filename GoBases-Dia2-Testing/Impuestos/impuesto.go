package main

import "fmt"

/*Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo
es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará
además un 10 % (27% en total).*/

func main() {
	Impuestos(510000)
}

func Impuestos(sueldo float32) int {
	var impuesto int
	switch {
	case sueldo > 50000:
		porcentaje := 0.83
		impuesto = int(sueldo - (sueldo * float32(porcentaje)))
		fmt.Println("El impuesto a pagar es de :", impuesto)
	case sueldo < 50000:
		fmt.Println("No debe pagar impuestos.")
	case sueldo > 150000 && sueldo > 50000:
		porcentaje := 0.73
		impuesto = int(sueldo - (sueldo * float32(porcentaje)))
		fmt.Println("El impuesto a pagar es de :", impuesto)
	}
	return impuesto
}
