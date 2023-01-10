package main

import "fmt"

func main() {
	Gustavo := Alumnos{
		Nombre:   "Gustavo",
		Apellido: "Aguero",
		DNI:      40725081,
		Fecha:    "21/10/1997",
	}
	Gustavo.Registrar(&alumnos)
	buscarAlumno := Alumnos{}
	buscarAlumno.Detalles("Gustavo")
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

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      uint
	Fecha    string
}

var alumnos = []Alumnos{}

func (alumno Alumnos) Registrar(storage *[]Alumnos) {
	*storage = append(*storage, alumno)
}

func (alumno Alumnos) Detalles(nombre string) {
	for _, alumno := range alumnos {
		if alumno.Nombre == nombre {
			fmt.Println("Nombre: ", alumno.Nombre, "\nApellido: ", alumno.Apellido, "\nDNI: ", alumno.DNI, "\nFecha: ", alumno.Fecha)
			break
		} else if alumno.Nombre != nombre {
			fmt.Println("No se encuentra el alumno")
		}
	}
}
