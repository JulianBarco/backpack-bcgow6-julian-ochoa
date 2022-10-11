package main

import (
	"fmt"
	"os"
)

// Setenv
/*
La función Setenv( key, value string ) modifica el valor de una variable de entorno recibiendo el nombre y el valor a asignar.
Retornará un error en caso de algún inconveniente.
*/


// Getenv
/*
La función Getenv( key string ) nos permitirá acceder a las variables de entorno del sistema, le pasaremos por parámetro la variable a la que deseamos acceder y nos devolverá su valor:
*/

// LookupEnv
/*
La función LookupEnv( key string ) es equivalente a la función Getenv( ) con la diferencia que retorna 2 valores:
El valor de la variable de entorno.
Un booleano para determinar si la variable existe o no.
*/

// ReadDir
/*
ReadDir(name string) ([]DirEntry, error) lee el directorio nombrado, devolviendo todas sus entradas de directorio ordenadas por nombre de archivo. Si se produce un error al leer el directorio, ReadDir devuelve las entradas que pudo leer antes del error, junto con el error.
*/

// ReadFile
/*
La función ReadFile( filename string ) recibe como parámetro la dirección y nombre del archivo en formato texto y nos devuelve el contenido del archivo en bytes o un error en caso que lo haya .
*/

// WriteFile
/*
La función WriteFile( filename string, data []byte, perm FileMode ) recibe como parámetro la dirección y nombre del archivo en formato texto, su contenido en formato bytes y el permiso que queramos asignarle.
Devuelve un error en caso que lo haya.
*/

func main(){

	err := os.Setenv("NAME","Julian")
	if err != nil{
		panic(err)
	}

	value := os.Getenv("NAME")
	fmt.Println(value)

	name, ok := os.LookupEnv("NAME")
	if !ok {
		panic("Variable no encontrada")
	}
	fmt.Println(name)

	flag := true

	if flag {
		fmt.Println(flag, "es verdadero")
	}else{
		fmt.Println(flag, "es falso")
		os.Exit(1)
	}
	fmt.Println("Fin de instrucciones")


	//ReadDir
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	fmt.Printf("coleted entries: %v\n", entries)

	//ReadFile
	// file, err := os.ReadFile("main.go")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(file))
	// fmt.Printf("coleted file: %v\n", file)

	//WriteFile
	// err = os.WriteFile("main.go", file, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	mensaje := "Primer mensaje por medio de writeFile"
	error := os.WriteFile(
		"./test.txt",
		[]byte(mensaje),
		0644)
	if error != nil{
		panic(error)
	}
}