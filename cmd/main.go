package main

import (
	"log"

	"github.com/facundo-vidiella/learning-go-web/cmd/handler"
	"github.com/facundo-vidiella/learning-go-web/pkg/store"

	"github.com/gin-gonic/gin"
)

func answerPong(c *gin.Context) {
	c.String(200, "pong")
}

func main() {
	var server = gin.Default()
	store.InitStore()

	server.GET("/ping", answerPong)

	server.GET("/products", handler.GetProducts)
	server.GET("/products/:id", handler.GetProductById)
	server.GET("products/search/:priceGt", handler.GetProductsByPrice)
	server.POST("/products", handler.CreateProduct)
	server.PUT("/products/:id", handler.UpdateProduct)
	server.PATCH("/products/:id", handler.UpdateProductField)
	server.DELETE("/products/:id", handler.DeleteProduct)

	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}
