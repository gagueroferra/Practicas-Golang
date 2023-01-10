package main

import (
	"errors"
	"fmt"
)

/*Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario es menor a 10.000"
y lanzalo en caso de que “salary” sea menor o igual a  10.000. La validación debe ser hecha con la función “Is()”
dentro del “main”.*/

func main() {

	coicidencia := errors.Is(Salary(9000), ThisError)
	fmt.Println(coicidencia)
}

var ThisError = errors.New("Error: el salario es menor a 10.000")
var ThisOtherError = errors.New("Otro error")

type MyError struct{}

func (err *MyError) Error() error {
	return ThisError
}

func Salary(salary int) error {
	if salary <= 10000 {
		myError := MyError{}
		return myError.Error()
	}
	return ThisOtherError
}