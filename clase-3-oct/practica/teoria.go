// package teoria

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// func teoria()  {
// 	fmt.Println("test")

// 	type productos struct{
// 		ID int
// 		Name string
// 		Color string
// 		Precio float64
// 		Stock	int
// 		Codigo string
// 		Publicado bool

// 	}

// 	nintedo := productos{
// 		ID: 1,
// 		Name: "Game boy",
// 		Color: "yellow",
// 		Precio: 499.9,
// 		Stock: 100,
// 		Codigo: "NGBOY1999",
// 		Publicado: true,
// 	}

// 	jsonData, err := json.Marshal(nintedo)
// 	if err != nil{
// 		log.Fatal("Error: ", err)
// 	}

// 	fmt.Printf("producto: \n %s \n", string(jsonData))

// }