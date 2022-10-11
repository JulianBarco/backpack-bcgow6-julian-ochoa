package main

import "fmt"

// Qué son los maps?
// Son estructuras que nos permiten crear variables tipo clave-valor definiendo un tipo de dato para las claves y otro para los valores

//Formas para declarar un map
//myMap := map[string]int{} nos permite definir el map con valores inicializados
//myMap := make(map[string]string) la segunda lo inicializa vacio

//Longitud de un map
//myMap := map[string]int{} nos permite definir el map con valores inicializados
// fmt.Prinln(len(myMap)) la función len devuelve cero para un map no inicializado.

func main() {

	//Inicializar un map y acceder a los valores
	var students = map[string]int{"julio": 20, "alec": 22}
	fmt.Println(students["alec"])
	fmt.Println("")
	fmt.Println("Agregar más valores al map")
	students["Lorenzo"] = 13
	students["Gina"] = 14
	fmt.Println("Map con valores agregados")
	fmt.Println(students)
	fmt.Println("")
	fmt.Println("Para actualizar valores")
	students["julio"]=22
	fmt.Println("Valor de Julio actualizado:",students["julio"])
	fmt.Println("")
	//Para recorrer los elementos de un map lo hacemos con un bucle for
	for llave, valor := range students{
		fmt.Println("llave:",llave,"==>","Valor",valor)
	}
	fmt.Println("")
	//Estraer datos de un map y asignarlos a una variable
	x, ok := students["julio"] //ok se implementa para saber si la llave está definida o no
	fmt.Println("Valor extraido de julio y guardado en X:", x)
	fmt.Println("x:",x,"ok:",ok)

	fmt.Println("")
	//Para ignorar la llave y obtener unicamente el valor, se usa _
	_, ok2 := students["julio"]
	fmt.Println("Ignorando la llave y obteniendo el valor",ok2)

	

}
