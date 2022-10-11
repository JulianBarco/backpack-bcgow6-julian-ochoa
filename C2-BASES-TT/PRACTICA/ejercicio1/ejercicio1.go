package main

import (
	"fmt"
)


type Alumno struct {
	Nombre string
	Apellido string
	DNI int
	FechaIngreso string
}

func (a Alumno) Detalles()  {
	fmt.Println("Nombre:",a.Nombre,"\nApellidos:",a.Apellido,"\nDNI:",a.DNI,"\nFecha de ingreso:",a.FechaIngreso)
}


func main() {


 a := Alumno{
	Nombre: "Julian",
	Apellido: "Barco",
	DNI: 12345678,
	FechaIngreso: "01/01/2020",
 }

 a.Detalles()



}