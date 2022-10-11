package main

import "fmt"

//Para los slices no hay necesidad de declara el tamaño, porque Go se
//encarga dinamicamente.

//Forma para declarar un arry
// var s []T  -->  var s = []bool{true,false}



func main(){

	var s = []bool{true,false}
	fmt.Println("Slices")
	fmt.Println("Slice S:", s)
	fmt.Println("")

	
	fmt.Println("declarar un slice apartir de un array")
	var myArr = []int{1,2,3,4,5}
	slc := myArr[0:2]
	fmt.Println(slc)
	fmt.Println("")
	fmt.Println("Para ver el tamaño de un slice")
	fmt.Println("Tamaño",len(slc)) //len va a devolver los elementos que tiene el slice 
	fmt.Println("Capacidad",cap(slc)) //La capacidad siempre va a devolver la cantidad de elementos que tiene el arry
	fmt.Println("slice antes del append",slc)
	fmt.Println("")
	//Usando append
	slc = append(slc, 1,2,3,4,5,6,7,8,9)
	fmt.Println("Tamaño luego del append",len(slc))
	fmt.Println("Capacidad luego del append",cap(slc))
	fmt.Println("slice despues del append",slc)

	//Otra sintaxis para declarar un slice (la más normal)
	fmt.Println("")
	var slc2 []int
	fmt.Println("segundo slice:",slc2)









}