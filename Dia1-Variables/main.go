package main

import "fmt"

func main() {
	//datos()
	weather()
}

/*Ejercicio 1 - Imprimí tu nombre
Tendrás que:
Crear una aplicación donde tengas como variable tu nombre y dirección.
Imprimir en consola el valor de cada una de las variables.*/

func datos() {
	var (
		nombre    = "Gustavo"
		direccion = "Pellegrini 658"
	)
	fmt.Println("Mi nombre es: ", nombre, "Mi direccion es: ", direccion)

}

/*Ejercicio 2 - Clima

Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.
Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprimí los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?*/

func weather() {
	temperatura := 35
	humedad := 15
	presion := 2
	fmt.Println("La temperatura es: ", temperatura, "\nLa humedad es del ", humedad, "%", "\nLa presion es de ", presion, "atm")
}
