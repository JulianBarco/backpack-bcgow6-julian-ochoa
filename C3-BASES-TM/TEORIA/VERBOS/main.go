package main


import "fmt"

/*
Verbos de impresion

%v    			valor en formato estándar
%T 				tipo de dato del valor a imprimir
%t   			bool
%s 			 	string
%f  			punto flotante
%d 				entero decimal
%b 				un entero binario
%o 				octal
%c 				imprime caracteres
%p 				dirección de memoria

*/

func main()  {


		//Practicando con un caracter de escape
		a := "Hola\n"
		b := "mundo"

		c,err := fmt.Print(a,b)
		if err != nil {
			panic(err.Error())
		}

		fmt.Print(c)


		flag := true

		fmt.Printf("La variable es %v",flag)

}