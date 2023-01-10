package main

import (
	"errors"
)

func main() {
	product1, err := NewProduct("large", 100) //creamos un nuevo objeto de la clase NewProduct con valores de tamaño y precio.
	if err != nil {                           //especificamos que pasa en caso de error
		panic(err)
	}
	println("El precio del producto es: ", int(product1.GetPrice())) //llamamos al metodo GetPrice() para implementarlo en nuestro objeto.
}

/*Ejercicio 2 - Productos
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga.*/

// crear estructura de cada uno de los casos
// crear instancias de cada estructura
// SmallProduct

type SmallSizeProduct struct {
	Price float64
}

func (product SmallSizeProduct) GetPrice() float64 {
	return product.Price
}

// MediumProduct
type MediumSizeProduct struct {
	Price float64
}

func (product MediumSizeProduct) GetPrice() float64 {
	return product.Price + (product.Price * 0.03)
}

// LargueProduct
type LargeSizeProduct struct {
	Price float64
}

func (product LargeSizeProduct) GetPrice() float64 {
	return product.Price + (product.Price * 0.06) + 2500
}

//creamos la interfaz

type Product interface {
	GetPrice() float64
}

//creamos un error nuevo para devolver

var InvalidProductType = errors.New("Invalid Product Type")

//creamos constantes con valores correspondientes a cada  instancia

const (
	SmallSizeProductType  = "small"
	MediumSizeProductType = "medium"
	LargeSizeProductType  = "large"
)

// creamos una funcion que devuelva la interfaz

func NewProduct(size string, price float64) (result Product, err error) {
	switch size {
	case SmallSizeProductType:
		result = SmallSizeProduct{Price: price}
	case MediumSizeProductType:
		result = MediumSizeProduct{Price: price}
	case LargeSizeProductType:
		result = LargeSizeProduct{Price: price}
	default:
		err = InvalidProductType
	}
	return
}
