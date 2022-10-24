package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//Copy
/*
La función io.Copy(dest Writer, orig Reader) toma un reader y copia los datos desde el origen al destino.
Devuelve el número de bytes copiados y el primer error encontrado durante la copia, si lo hubiera.
*/

//ReadAll
/*
ReadAll(r Reader) lee desde r hasta un error o EOF y devuelve los datos que leyó y un error si lo hubiera. 
*/

//WriteString
/*
La función WriteString(w Writer, s string) que toma una cadena de texto y un Writer, escribe el contenido de s a w, que acepta una slice de bytes.
*/

func main() {
	reader := strings.NewReader("Hola mundo\n")

	_,err := io.Copy(os.Stdout,reader)
	if err != nil {
		panic(err)
	}

	r := strings.NewReader("some io.Reader stream to be read")

	b,err := io.ReadAll(r);
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	//WriteString
	n,err := io.WriteString(os.Stdout,"Hola mundo\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(&n) // direccion de memoria
	fmt.Println(n) // valor de la variable


}