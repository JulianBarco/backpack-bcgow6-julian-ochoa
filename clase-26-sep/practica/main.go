package main

import "fmt"

/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
*/

func impuesto(sueldo float64) float64 {
	if sueldo >= 50.0000 && sueldo < 150.000 {
		descuento := sueldo - (sueldo * 0.17)
		return float64(descuento)
	} else {
		descuento := sueldo - (sueldo * 0.10)
		return float64(descuento)

	}

}


//Ejercicio 2 - Calcular promedio

/*
Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/

func promedio(notas ... float64)(prom float64){
	
	var sumatoriaNotas float64
	
	count := 0

	for _, notas := range notas{
		sumatoriaNotas += notas
		count += count 
	}
	prom = sumatoriaNotas / float64(len(notas))

	return 

}



//Ejercicio 3 - Calcular salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.


func sueldoPorHoras(horas float64, categoria string){

	switch categoria {
	case "A":
		salario := 3000 * horas
		salarioIncremento := salario + (salario * 0.5)
		fmt.Println("El salario es", salarioIncremento)
	case "B":
		salario := 1500 * horas
		salarioIncremento := salario + (salario * 0.2)
		fmt.Println("El salario es", salarioIncremento)
	case "C":
		salario := 1000 * horas
		fmt.Println("El salario es", salario)

	}
}


/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
Ejemplo:
*/


const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )

func calMin(valores ... float64)(result float64){
	result = valores[0]
	for _, num := range valores{

		if num < result{
			result = num
		}
	}
	return 
}

func calMax(valores ... float64)(result float64)  {
	
	for _, num := range valores{
		

		if num > result{
			
			result = num
		}
	}
	return 
}

func calAverage(valores ... float64)(prom float64) {
	var sumValores float64 = 0
	var cont float64 = 0 

	for _, num := range valores{
		sumValores += num
		cont += 1
	}

	prom = sumValores/ cont

	return prom
}


func funcOperadora(caso string)func(...float64)float64{
	switch caso{
	case minimum:
		return calMin
	case average:
		return calAverage
	case maximum:
		return calMax
	}
	return nil
}



func main() {

	//######### Ejecución primer ejercicio #########
	ejercicio1 := impuesto(150.000)
	fmt.Println("######### Solucion tercer punto #########")
	fmt.Println("El sueldo más el descuento es de:", ejercicio1)
	//######### final primer ejercicio #########
	fmt.Println("")
	//######### Ejecución segundo ejercicio #########
	fmt.Println(promedio(4,5,3,4,5,6))
	//######### final segundo ejercicio #########
	fmt.Println("")

	//######### Ejecución tercer ejercicio #########
	fmt.Println("######### Solucion tercer punto #########")
	sueldoPorHoras(10, "C")
	//######### final tercer ejercicio #########
	fmt.Println("")

	//######### Ejecución cuarto ejercicio #########
	calMax(1,2,3,4)
	calMin(1,2,3,4)
	calAverage(1,2,3,4)
	// fmt.Println(funcOperadora("average"))
	funcOperadora("minimum")
}
