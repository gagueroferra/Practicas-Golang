package main

import "fmt"

func main() {
	Promedio(2, 4, 6, 8, 10)
}

/*Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.*/

func Promedio(notas ...int) float32 {
	var suma int
	for i := 0; i < len(notas); i++ {
		suma += notas[i]
	}
	promedio := suma / len(notas)
	fmt.Println("El promedio es de: ", promedio)
	return float32(promedio)
}
