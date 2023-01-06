package main

import "fmt"

func main() {
	//impuestos(51000)
	//promedioXestudiante(2, 2, 2, 2, 2)
	calcular(10000, "A")
}

/*Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo
es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará
además un 10 % (27% en total).*/

func impuestos(sueldo float64) float64 {
	var impuesto float64
	switch {
	case sueldo > 50000:
		porcentaje := 0.83
		impuesto = sueldo - (sueldo * porcentaje)
		fmt.Printf("El impuesto a pagar es de: %.2f\n", impuesto)
	}
	return impuesto
}

/*Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se
le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.*/

func promedioXestudiante(notas ...int) int {
	var suma int
	for i := 0; i < len(notas); i++ {
		suma += notas[i]
	}
	prom := suma / len(notas)
	fmt.Println("El promedio de las notas es: ", prom)
	return prom
}

/*Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.*/

func calcular(minutos int, categoria string) int {
	var sueldo_porcentaje int
	minutos_a_hora := minutos / 60
	fmt.Println("La cantidad de minutos equivale a: ", minutos_a_hora, "horas mensuales.")
	switch {
	case categoria == "A":
		salarioXh := 3000
		porcentaje := 0.50
		sueldo := minutos_a_hora * salarioXh
		sueldo_porcentaje = sueldo + (sueldo * int(porcentaje))
	case categoria == "B":
		salarioXh := 1500
		porcentaje := 0.20
		sueldo := minutos_a_hora * salarioXh
		sueldo_porcentaje = sueldo + (sueldo * int(porcentaje))
	case categoria == "C":
		salarioXh := 1000
		sueldo_porcentaje = minutos_a_hora * salarioXh
	}
	fmt.Println("El sueldo mensual segun minutos y categoria especificada es: ", sueldo_porcentaje)
	return sueldo_porcentaje
}

/*Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.
Ejemplo:

const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)

...

minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...*/
