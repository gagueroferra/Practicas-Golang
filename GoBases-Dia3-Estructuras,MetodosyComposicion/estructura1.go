package main

import (
	"fmt"
)

/*Ejercicio 1
Crear un programa que cumpla los siguiente puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto
desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().*/

type Product struct {
	ID          uint
	Name        string
	Price       float32
	Description string
	Category    string
}

var products = []Product{
	{ID: 1, Name: "Pelota", Price: 1000, Description: "Pelota Adidas", Category: "Futbol"},
	{ID: 2, Name: "Guinda", Price: 3000, Description: "Guinda Rugby", Category: "Rugby"},
}

func (product Product) Save(storage *[]Product) {
	*storage = append(*storage, product)
	fmt.Println(products)
}
func (product Product) GetAll() []Product {
	fmt.Println(products)
	return products
}
func (product Product) GetById(id int) Product {
	for _, product := range products {
		if id == int(product.ID) {
			fmt.Println("El producto es: ", product)
		}
	}
	return product
}

func main() {
	newProduct := Product{
		ID:          3,
		Name:        "Pelotita",
		Price:       200,
		Description: "Pelotita de tenis",
		Category:    "Tenis",
	}
	newProduct.Save(&products)
	producto := newProduct
	producto.GetById(3)
}
