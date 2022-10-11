package main

import "fmt"

// Estructura en go
// El concepto fundamental se llama composición, lo que significa que una estructura puede contener otra estructura, y así sucesivamente.
type Perro struct {
	Nombre      string
	Genero      string
	Edad        int
	Raza        string
	Peso        float64
	Propietario Dueño
}

type Dueño struct {
	Nombre           string
	Edad             int
	NumeroDeTelefono int
}

func main() {

	// Existen dos formas de instaciar una estructura
	// Si el nombres está en mayúscula es publico
	p1 := Perro{
		Nombre: "Dogo",
		Genero: "Macho",
		Edad:   22,
		Raza:   "pitbull",
		Peso:   70.5,
	}

	//Ejemplo de como inizializar una estructura aniadida
	/*
		p1.Propietario = Dueño{
			Nombre: "Dogo",
			Genero: "Macho",
			Edad:   22,
			Raza:   "pitbull",
			Peso:   70.5,
				Propietario: Dueño{
					Nombre:           "Juan",
					Edad:             22,
					NumeroDeTelefono: 123456789,
				}
		}
	*/

	fmt.Printf("Persona %+v", p1)
	fmt.Println("")
	// para actualizar un valor de la estructura
	p1.Nombre = "Julio Cesar"
	// fmt.Printf(p1.Nombre)

	// Para asignarle un valor a una estructura dentro de otra
	p1.Propietario = Dueño{
		Nombre:           "Julio Cesar",
		Edad:             22,
		NumeroDeTelefono: 123456789,
	}
	fmt.Println("")
	fmt.Printf("Persona %+v", p1)
	fmt.Printf("Propietario %+v", p1.Propietario)

}
