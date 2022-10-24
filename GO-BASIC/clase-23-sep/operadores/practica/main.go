package main

import "fmt"

func main() {

	fmt.Print("########## Primer punto ##########")
	palabra := "hola"

	fmt.Println(len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Println(i, "letra", string(palabra[i]))
	}
	fmt.Println("########## fin primer punto ##########")
	fmt.Println("")

	fmt.Println("########## Segundo punto ##########")

	/*Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
	Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
	Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.*/

	type user struct {
		edad    int
		nombre  string
		trabaja bool
		años    float64
		sueldo  float64
	}

	cliente := user{
		edad:    22,
		nombre:  "Julian Barco",
		trabaja: false,
		años:    1,
		sueldo:  200.000,
	}

	fmt.Println(!(cliente.trabaja))

	if cliente.edad < 20 {
		fmt.Println("El cliente no tiene la edad suficiente")
	}
	if !cliente.trabaja {
		fmt.Println("El cliente no tiene se encuentra trabajando")

	}
	if cliente.años <= 1 {
		fmt.Println("El cliente lleva muy poco en la compañia")

	}
	if cliente.sueldo > 100.000 {
		fmt.Println("Al cliente no se le cobrarían intereses")
	} else {
		fmt.Println("Al cliente se le cobran intereses")
	}
	if cliente.edad > 20 && cliente.trabaja && cliente.sueldo > 100.000 && cliente.años > 1{
		fmt.Println("El cliente puede acceder al credito")
	} else {
		fmt.Println("El cliente no cumple con todos los requisitos para el prestamo")
	}
	fmt.Println("########## fin segundo punto ##########")
	fmt.Println("")

	fmt.Println("########## Tercer punto ##########")

	/*
	Ejercicio 3 - A qué mes corresponde

	Realizar una aplicación que contenga una variable con el número del mes. 
	Según el número, imprimir el mes que corresponda en texto. 
	¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
	Ej: 7, Julio
	*/

	numero := 33

	switch numero{
	case 1:
		fmt.Println("El",numero,"corresponde a enero")
	
	case 2:
		fmt.Println("El",numero,"corresponde a febrero")
	
	case 3:
		fmt.Println("El",numero,"corresponde a marzo")

	default:
		fmt.Println("Ingresa un numero entre 1 - 12")

	}
	fmt.Println("########## fin tercer punto ##########")
	fmt.Println("")

	fmt.Println("########## Cuarto punto ##########")

	/*
	Ejercicio 4 - Qué edad tiene...
	Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin. 
	*/
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	

	fmt.Println("La edad de","Benjamin","es",employees["Benjamin"])
	count := 0
	for llave, valor := range employees{
		fmt.Println(llave)

		if valor > 21{
			count ++
		}
		
	}

	println("la cantidad de empleados mayores a 21 años es", count)
	
	//Eliminando a pedro del mapa
	delete(employees,"Pedro")

	fmt.Println("Mapa sin pedro", employees)
}
