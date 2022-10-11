package main

//En go solo exite un tipo de bucle, y es el for
//Hay diferentes formas de utilizarlos, estas son:
// Standar For - For Range - Bucle infinito - Bucle While

//######## El for standar ########
// for i := 0; i < 100; i++{
// 	sum += i
// }

//######## El for range ########
// frutas := []string{"manzana","pera","guayaba"}
// for i, frutas := range frutas{
//	fmt.Println(i,frutas)
//}

//######## Bucle infinito ########
// sum := 0
// for{
// 	sum++
//}
//######## Bucle while ########
// sum := 1
// for sum < 10 {
// 	sum += sum
//}
//	fmt.Println(sum)

// Para romper un bucle utilizamos la palabra reservada break
// sum := 0
//for{
// 	sum++
//  if sum >= 1000{
// 	 break
//  }
//}
//fmt.Println(sum)

//También existe la palabra continue, esta lo que hace es continuar con la siguiente interación.
/*

	for i := 0, i < 10; i++{
		if i%2 == 0{
			continue
		}
		fmt.Println(i,"is odd")
	}



*/

func main() {

}
