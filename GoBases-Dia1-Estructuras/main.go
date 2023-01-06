package main

import "fmt"

func main() {
	//letras("Hola")
	//prestamo(25, "desempleado", 2, 50000)
	//month(5)
	edades("Benjamin")
}

/*Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
Luego, imprimí cada una de las letras.*/

func letras(palabra string) {
	fmt.Println("La palabra ", palabra, "tiene", len(palabra), "letras.")
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
	}
}

/*Ejercicio 2 - Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente
se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.
Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.*/

func prestamo(edad uint, situacion string, antiguedad uint, sueldo float64) {
	switch {
	case edad > 21 && situacion == "empleado" && antiguedad > 1:
		fmt.Println("Edad: ", edad, "\nSituacion: ", situacion, "\nAntiguedad: ", antiguedad, "\nEl prestamo ha sido aprobado.")
		if sueldo > 100000 {
			fmt.Println("No se le cobrara intereses.")
		} else if sueldo < 100000 {
			fmt.Println("Se le cobrara intereses.")
		}
	default:
		fmt.Println("No cumple los requisitos para acceder al prestamo.")
	}
}

/*Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que reciba  una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
Nota: Validar que el número del mes, sea correcto.*/

func month(numero uint) {
	meses := map[uint]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}
	fmt.Println("El mes que corresponde al numero ", numero, "pertenece al mes: ", meses[numero])
}

/*Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.*/

func edades(nombre string) {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("El empleado ", nombre, " tiene ", employees[nombre], "años.")
}
