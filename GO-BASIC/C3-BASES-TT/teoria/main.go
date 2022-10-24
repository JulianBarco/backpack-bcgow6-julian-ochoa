package main

import "fmt"




func main(){
	// Formas para declarar punteros
	// var puntero *int
	// p2 := new(int)

	var numero int = 10
	p3 := &numero

	fmt.Printf("La direccion es: %p\n", &numero)
	fmt.Printf("La direccion es: %p\n", &p3)

	// El operardor de desreferencia es *
	// sirve para obtener el valor de una variable atraves de su puntero

	fmt.Printf("El valor de numero es: %d\n", numero)
	*p3 = 20
	fmt.Printf("El valor de numero es: %d\n", numero)

	fmt.Println("Usando la funcion Incrementar")

	Incrementar(p3)
	fmt.Printf("El valor de numero sin la funcion es: %d\n", numero)
	fmt.Printf("El valor de numero usando la funcion es: %d\n", numero)

	fmt.Println("Usando el metodo Actualizar")

	var c = Coso{
		valor: 10,
	}

	fmt.Printf("El valor de c antes era: %+v\n", c.valor)

	c.Actualizar(20)

	fmt.Printf("El valor de c despues es: %+v\n", c.valor)



}

func Incrementar(puntero *int){
	*puntero = 40
}


type Coso struct{
	valor int
}

func (c *Coso)Actualizar(new int){
	c.valor = new
}