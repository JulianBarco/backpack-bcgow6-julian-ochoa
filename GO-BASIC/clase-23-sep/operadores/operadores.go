package main

import "fmt"

func main() {

	x, y := 10, 20
	fmt.Println(x % y)
	x++
	y++
	fmt.Println("X incrementada", x)
	fmt.Println("Y incrementada", y)
}
