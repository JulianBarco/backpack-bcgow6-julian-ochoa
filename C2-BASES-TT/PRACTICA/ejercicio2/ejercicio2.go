package main

import "fmt"

/*

Ejercicio 3 - Productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

*/

func main() {

	// Crear una tienda
	// type Tienda struct {
	// 	Productos []Producto
	// }

	type Producto struct{
		Tipo string
		Nombre string
		Precio float64
	}

	// type IProducto interface{
	// 	CalcularCosto() float64
	// }

	// type Ecommerce interface{
	// 	Total() float64
	// 	Agregar(producto Producto) error
	// }

	func (p *Producto) NuevoProducto(tipo, nombre string, precio float64) *Producto {
		return &Producto{
			Tipo:   tipo,
			Nombre: nombre,
			Precio: precio,
		}
	}

	// func nuevaTienda() Ecommerce{
	// 	return &Tienda{}
	// }

	// func (p *Producto) CalcularCosto() float64{
	// 	switch p.Tipo{
	// 	case "mediano":
	// 	return p.Precio + ( p.Precio * 0.03)

	// 	case "grande":
	// 	return p.Precio + ( p.Precio * 0.03) + 2500
		
	// 	}
	// 	return p.Precio
	// }


	// func (t *Tienda) Total() float64 {
	// 	var total float64

	// 	for _,total := range t.Productos{
	// 		total += producto.CalcularCosto()
	// 	}

	// 	return total
	// }
}
