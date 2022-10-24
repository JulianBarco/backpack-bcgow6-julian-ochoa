package main

import "fmt"

//los arrays no se les puede modificar el tamaño 
//Forma para declarar un arry
// var prueba [2]string
//Forma para asignarle un valor en determinada posición
//a[0] = "Hello"
//b[1] = "World"
//Forma para declarar arrys  con valores predeterminados
//var myArr = [...]int{1,2,3} 


func main() {
	var prueba [2]string
	prueba[0] = "Hello"
	prueba[1] = "World"
	
	fmt.Println("Arryas y slices")
	fmt.Println("Posición 0",prueba[0])
	fmt.Println("Posición 1",prueba[1])

	fmt.Println("Array con valores predeterminados")
	// fmt.Println("")
	var myArr = [...]int{1,2,3}
	fmt.Println(myArr)

}
