package main

import (
	"fmt"
)

/*Ejercicio 4 - Impuestos de salario #4
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el
valor de “salary”
indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es de
150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).*/

func main() {
	salario := Salary(160000)
	fmt.Println(salario)
}

func Salary(salary int) error { //creamos una funcion que pida un salario por parametro y devuelva un error.
	if salary <= 150000 {
		err := fmt.Errorf("%s %d", "Error: el minimo imponible es de 150.000 y el salario ingresado es de: ", salary)
		//utilizamos el paquete errorf para crear un error con 2 parametros. 1) un string, 2) un dato int o float.
		return err
	}
	err2 := fmt.Errorf("%s %d", "El sueldo pasa el minimo imponible. El sueldo ingresado es de: ", salary)
	return err2
}
