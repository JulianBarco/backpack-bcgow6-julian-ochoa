package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Julian",
		})
	})

	router.Run()
}