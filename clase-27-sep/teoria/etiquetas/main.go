package main

import (
	"encoding/json"
	"fmt"
)

/*
Dentro de nuestras estructuras podemos definir etiquetas o anotaciones que hagan referencia a cada uno de los campos que aparecen luego de declarar el tipo de dato.
*/

type Dueño struct {
	Nombre           string
	Edad             int
	NumeroDeTelefono int
}

type Perro struct {
	Nombre      string  `json:"nombre"`
	Genero      string  `json:"genero"`
	Edad        int     `json:"edad"`
	Raza        string  `json:"raza"`
	Peso        float64 `json:"peso"`
	Propietario Dueño   `json:"propietario"`
}

func main() {
	//Se inicializa la estructura Perro
	p1 := Perro{
		Nombre: "Dogo",
		Genero: "Macho",
		Edad:   22,
		Raza:   "pitbull",
		Peso:   70.5,
		Propietario: Dueño{
			Nombre:           "Julian",
			Edad:             22,
			NumeroDeTelefono: 123456789,
		},
	}

	fmt.Printf("Persona %+v", p1)
	fmt.Println("")
	b,err := json.Marshal(&p1)
	if err != nil{
		panic(err)
	}
	fmt.Println("")
	fmt.Println(string(b))

}
