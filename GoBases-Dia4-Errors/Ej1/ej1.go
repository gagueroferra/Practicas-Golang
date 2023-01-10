package main

import (
	"errors"
	"fmt"
)

/* Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario ingresado
no alcanza el mínimo imponible"
y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el mensaje
“Debe pagar impuesto”.*/

func main() {
	Salary(160000)
}

type MyError struct{}

func (err *MyError) Error() error {
	return myError
}

var myError = errors.New("Error: el salario ingresado no alcanza el mínimo imponible")

func Salary(salary int) error {
	if salary < 150000 {
		myErr := MyError{}
		fmt.Println(myErr.Error())
	} else {
		fmt.Println("Debe pagar impuestos")
	}
	return nil
}
