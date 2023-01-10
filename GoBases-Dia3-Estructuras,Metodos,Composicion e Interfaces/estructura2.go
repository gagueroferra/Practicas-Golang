package main

import "fmt"

/*Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a
gestionar correctamente dichos empleados. Los objetivos son:
Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos
de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método
PrintEmployee().
Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.*/

type Person struct {
	ID          uint
	Name        string
	DateOfBirth string
}
type Employee struct {
	ID       uint
	Position string
	Data     Person
}

var employee = []Employee{}

func (employee Employee) PrintEmployee() {
	fmt.Printf("%+v", employee)
}

func (person Person) PrintPerson() {
	fmt.Printf("%+v", person)
}

func main() {
	newEmployee := Employee{
		ID:       1,
		Position: "Software Developer",
		Data:     Person{ID: 40725081, Name: "Gustavo Aguero", DateOfBirth: "21/10/1997"},
	}
	newEmployee.PrintEmployee()

	newPerson := Person{
		ID:          13144682,
		Name:        "Roberto Aguero",
		DateOfBirth: "13/10/1959",
	}
	newPerson.PrintPerson()
}
