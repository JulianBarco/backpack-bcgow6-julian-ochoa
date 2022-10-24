package main

import (
	"fmt"
	"os"
)

/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

type producto struct {
	id       int
	precio   float32
	cantidad int
}

func GuardarArchivo(prod producto, file os.File) (err error) {
	_, errG := file.WriteString(fmt.Sprint(
		prod.id, ";",
		prod.precio, ";",
		prod.cantidad, "\n"))

	if errG != nil {
		return errG
	}
	return
}

func main() {

	productosEnArchivo, err := os.Create("infoProductos.txt")
	if err != nil {
		panic("No se pudo crear el archivo")
	}

	p1 := producto{
		id: 1,
		precio: 1000,
		cantidad: 100,
	}
	p2 := producto{
		id: 2,
		precio: 5000,
		cantidad: 700,
	}
	p3 := producto{
		id: 3,
		precio: 111,
		cantidad: 111,
	}
	GuardarArchivo(p1,*productosEnArchivo)
	GuardarArchivo(p2,*productosEnArchivo)
	GuardarArchivo(p3,*productosEnArchivo)
}
