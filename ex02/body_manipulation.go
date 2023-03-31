package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `json:"nombre"`
	LastName string `json:"apellido"`
}

func main() {
	router := gin.Default()
	var person Person
	router.POST("/saludo", func(ctx *gin.Context) {
		if err := ctx.BindJSON(&person); err != nil {
			log.Fatal(err)
		}
		message := "Hola " + person.Name + " " + person.LastName

		ctx.Data(200, "text/html; charset=utf-8", []byte(message))
	})

	router.Run()
}
