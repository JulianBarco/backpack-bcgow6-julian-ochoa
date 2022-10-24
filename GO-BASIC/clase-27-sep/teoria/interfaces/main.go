package main

import "fmt"

//Qués es una interfaz?
//Una interfaz es una forma de definir métodos que deben ser utilizados pero sin definirlos.

//Para que se utilizan las interfaces?
//Las interfaces son utilizadas para brindar modularidad al lenguaje.

type Gato struct {
	Nombre      string
	Genero      string
	Edad        int
	Raza        string
	Peso        float64
}

type Perro struct {
	Nombre string
	Genero string
	Edad   int
	Raza   string
	Peso   float64
}

func (p Perro) Dormir() {
	fmt.Println(p.Nombre,"esta durmiendo")
}

func (p Perro) Comer() {
	fmt.Println("soy", p.Nombre ,"y estoy comiendo")
}

func (g Gato) Dormir() {
	fmt.Println("soy", g.Nombre ,"y estoy comiendo")
}

func (g Gato) Comer() {
	fmt.Println(g.Nombre,"esta durmiendo")
}

type Animal interface {
	Dormir()
	Comer()
}

// Definimos la nueva estructura
func NewPerro() Animal {
	//Debo devolver una referencia a la estructura que está intentando implementar la interfaz
	return &Perro{

	}
}

/*
Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/


type Matrix struct{
	
}

func main() {

	//############# Solucion primer punto #############

	dog := Perro{
		Nombre: "Gina",
		Genero: "Hembra",
		Edad:   22,
		Raza:   "pitbull",
		Peso:   70.5,
	}

	cat := Gato{
		Nombre: "Amarrillo",
		Genero: "Macho",
		Edad:   22,
		Raza:   "pitbull",
		Peso:   70.5,
	}

	dog.Dormir()
	dog.Comer()

	cat.Dormir()
	cat.Comer()
	
	fmt.Println("")
	a := NewPerro()
	a.Comer()
	a.Dormir()

	//############# Final primer punto #############

	//############# Solucion primer punto #############


	

}


