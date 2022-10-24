package practica

import "fmt"

// 1 punto
// Ejercicio 1 - Imprimí tu nombre
// Crea una aplicación donde tengas como variable tu nombre y dirección.
// Imprime en consola el valor de cada una de las variables.

// 2 punto
// Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares.
// Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
// Imprime los valores de las variables en consola.
// ¿Qué tipo de dato le asignarías a las variables?

var temperatura int
var humedad float64
var presión int

// tercer punto
// var 1nombre string incorrecta
// var apellido string correcta
// var int edad incorrecta
// 1apellido := 6 incorrecta
// var licencia_de_conducir = true correcta
// var estatura de la persona int incorrecta
// cantidadDeHijos := 2 correcta siempre y cuando este en una función.

//Cuarto punto
// Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:
// Verificar su código y realizar las correcciones necesarias.

var (
	nombre    = "Julian Ochoa Barco"
	dirección = "Cra 90c # 15 a 19 Medellín Colombia"
)

func practica() {

	//Solución ejercicio 1
	fmt.Println("Primer punto")
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Nombre: ", dirección)
	fmt.Println("")

	//Solución ejercico 2
	temperatura = 18
	humedad = 98
	presión = 1026
	fmt.Println("Segundo punto")
	fmt.Println("temperatura:", temperatura, "C˚")
	fmt.Println("Humedad:", humedad, "%")
	fmt.Println("Presión:", presión, "hPa")

	fmt.Println("")
	fmt.Println("Cuarto punto")
	var apellido string = "Gomez"
	//Original
	// var edad int = "35"
	//Corregida
	var edad int = 35
	//Original
	//boolean := false
	//Corregida
	var boolean bool
	boolean = true
	//Original
	// var sueldo string = 45857.90
	//Corregida
	var sueldo float64 = 45857.90
	var nombre string = "Julián"
	fmt.Println("Testeo y correción de variables")
	fmt.Println("La primera está perfecta: ", apellido)
	fmt.Println("a La segunda se le quitan las \"\": ", edad)
	fmt.Println("a La tercera se le la estructura correcta: ", boolean)
	fmt.Println("a La cuarta se le quitan las \"\":: ", sueldo)
	fmt.Println("La ultima está perfecta: ", nombre)
}
