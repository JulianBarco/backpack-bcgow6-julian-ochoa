package main

import (
	"fmt"
	"os"
)

// La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

func main()  {
	
	file, err := os.ReadFile("infoProductos.txt")
	if err != nil {
		panic("Error leyendo el archivo")
	}
	fmt.Println(string(file))

}