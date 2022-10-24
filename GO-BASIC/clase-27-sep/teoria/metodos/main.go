package main

import (
	"fmt"
)

/*
Go no tiene clases. Sin embargo, se pueden definir métodos en los tipos de datos.
Un método es una función con un argumento de receptor especial. El receptor aparece en su propia lista de argumentos entre la palabra clave func y el nombre del método.
*/

type Perro struct {
	Nombre      string
	Genero      string
	Edad        int
	Raza        string
	Peso        float64
}

// Ahora haremos que el perro ladre

//La letra p cumple la función de receptor, es decir, es el que recibe el valor de la estructura Perro
//La p se puede comprar con el this de java, o el self de python

func (p Perro) Ladrar(){
	fmt.Println("Guau Guau, soy", p.Nombre)
}

func (p Perro) Renombrar(new string)(string){
	p.Nombre = new
	return "Guau Guau,ahora soy" + p.Nombre
}

func main()  {
	

	dog := Perro{
		Nombre: "Dogo",
		Genero: "Macho",
		Edad:   22,
		Raza:   "pitbull",
		Peso:   70.5,
	}

	dog.Ladrar()
	dog.Renombrar("Ginita")

	// Acá solo se está trabajando sobre una copia de la estructura orginal, por lo que no se modifica el valor de la estructura original. Para modificar el valor de la estructura original se debe usar punteros(*)

}