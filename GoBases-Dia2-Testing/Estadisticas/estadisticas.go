package main

import (
	"errors"
	"fmt"
)

func main() {
	calculadora, err := Estadistica(Promedia)
	if err != nil {
		panic(err)
	}
	calculadora(10, 9, 5, 2, 5, 1, 6, 3)
}

/*Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso.
Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y
un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.*/

func Estadistica(operacion string, notas ...uint) (operation func(...float32) float32, err error) {
	switch operacion {
	case Minima:
		operation = Minimum
	case Maxima:
		operation = Maximum
	case Promedia:
		operation = Average
	}
	return
}

const (
	Minima   = "minimo"
	Maxima   = "maximo"
	Promedia = "promedio"
)

var (
	ErrorOperadorNoEncontrado = errors.New("el operador es incorrecto")
)

func Minimum(notas ...float32) float32 {
	var aux float32
	for i := 0; i < len(notas); i++ {
		for j := 0; j < len(notas); j++ {
			if notas[i] < notas[j] {
				aux = notas[i]
				notas[i] = notas[j]
				notas[j] = aux
			}
		}
	}
	fmt.Println("La nota menor es: ", notas[0])
	return notas[0]
}

func Maximum(notas ...float32) float32 {
	var aux float32
	for i := 0; i < len(notas); i++ {
		for j := 0; j < len(notas); j++ {
			if notas[i] > notas[j] {
				aux = notas[i]
				notas[i] = notas[j]
				notas[j] = aux
			}
		}
	}
	fmt.Println("La nota mayor es: ", notas[0])
	return notas[0]
}

func Average(notas ...float32) float32 {
	var suma float32
	var promedio float32
	for i := 0; i < len(notas); i++ {
		suma += float32(notas[i])
	}
	promedio = suma / float32(len(notas))
	fmt.Println("El promedio de las notas es de: ", promedio)
	return promedio
}
