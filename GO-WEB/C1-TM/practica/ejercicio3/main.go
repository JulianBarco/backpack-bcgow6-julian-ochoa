package main

import (
	"encoding/json"
	"fmt"
	"os"

	//"time"

	"github.com/gin-gonic/gin"
)

const jsonRute = "/Users/julochoa/Documents/backpack-bcgow6-julian-ochoa/GO-WEB/C1-TM/practica/ejercicio1/products.json"

type Products struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Colour       string  `json:"colour"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         int     `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creation_date"`
}

func (p Products) getFromFile(filePath string) (*[]Products, error) {
	var products []Products

	rawData, err := os.ReadFile(filePath)

	if err != nil {
		return &products, err
	}

	

	err = json.Unmarshal(rawData, &products)

	if err != nil {
		return &products, err
	}

	return &products, nil
}

func getAll(ctx *gin.Context) {
	products, err := Products{}.getFromFile(jsonRute)
	fmt.Println(jsonRute)
	fmt.Println(products)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, products)
}

func Home(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Bienvenido",
	})
}

func main() {
	fmt.Println(Products{})

	router := gin.Default()

	router.GET("/data", Home)
	router.GET("/alldata", getAll)
	router.Run()

}
