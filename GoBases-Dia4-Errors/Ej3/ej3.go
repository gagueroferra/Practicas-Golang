package main

import (
	"errors"
	"fmt"
)

/*Ejercicio 3 - Impuestos de salario #3

Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()”.*/

func main() {
	coincidencia := errors.Is(Salary(10000), ThisError)
	fmt.Println(coincidencia)
}

var (
	ThisError      = errors.New("error: el salario es menor a 10.000")
	ThisOtherError = errors.New("error general")
)

func Salary(salary int) error {
	if salary <= 10000 {
		return ThisError
	}
	return ThisOtherError
}
