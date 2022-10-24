package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

const jsonRute = "../json/products.json"

type Product struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Colour       string    `json:"colour"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	Code         int       `json:"code"`
	Published    bool      `json:"published"`
	CreationDate string `json:"creation_date"`
}

func (p Product) getFromFile(filePath string) (*[]Product, error) {
	var products []Product

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
	products, err := Product{}.getFromFile(jsonRute)
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

func getOne(ctx *gin.Context){
	var product Product

	products, err := Product{}.getFromFile(jsonRute)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, item := range *products {
		if item.Id == idInt {
			product = item
		}
	}

	ctx.JSON(200, product)
}


func main() {
	fmt.Println(Product{})

	router := gin.Default()

	router.GET("/", Home)
	router.GET("/data", getAll)
	router.GET("/data/:id", getOne)
	router.Run()

}
