package main

import (
	"errors"
	"fmt"
)

func main() {
	newStudent := Student{
		Name:    "Gustavo",
		Surname: "Aguero",
		DNI:     40725081,
		Date:    "21/10/1997",
	}
	newStudent.RegisterStudent(&students)
	newStudent.GetStudentData("Gustav")
}

/*
Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de
cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle.*/

type Student struct { //creamos la estructura
	Name    string
	Surname string
	DNI     uint
	Date    string
}

var students = []Student{}

var err = errors.New("Error: Student not found") //creamos un error para devolver

func (student Student) RegisterStudent(storage *[]Student) { //instanciamos la estructura
	*storage = append(*storage, student)
}

func (student Student) GetStudentData(name string) { //instanciamos la estructura
	switch name {
	case student.Name:
		fmt.Println("Name: ", student.Name, "\nSurname: ", student.Surname, "\nDNI: ", student.DNI, "\nDate: ", student.Date)
	default:
		fmt.Println(err)
	}
}
