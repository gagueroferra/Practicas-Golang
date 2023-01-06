package main

import "fmt"

func main() {
	Salario(10000, "A")
}

/*Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas t
rabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes,
la categoría y que devuelva su salario.*/

func Salario(minutosXmes int, categoria string) float32 {
	mAh := minutosXmes / 60
	var salarioXmes float32
	switch {
	case categoria == "A":
		sXh := 3000
		porcentaje := 0.50
		salario := float32(mAh) * float32(sXh)
		salarioXmes = salario + (salario * float32(porcentaje))
		fmt.Println("La cantidad de horas trabajadas por mes es: ", mAh, "\nLa categoria es: ", categoria, "\nEl salario es de: ", salarioXmes)
	case categoria == "B":
		sXh := 1500
		porcentaje := 0.80
		salario := float32(mAh) * float32(sXh)
		salarioXmes = salario + (salario * float32(porcentaje))
		fmt.Println("La cantidad de horas trabajadas por mes es: ", mAh, "\nLa categoria es: ", categoria, "\nEl salario es de: ", salarioXmes)
	case categoria == "C":
		sXh := 1000
		salario := float32(mAh) * float32(sXh)
		salarioXmes = salario
		fmt.Println("La cantidad de horas trabajadas por mes es: ", mAh, "\nLa categoria es: ", categoria, "\nEl salario es de: ", salarioXmes)
	}
	return salarioXmes
}
