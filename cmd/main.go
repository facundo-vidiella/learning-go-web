package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/facundo-vidiella/learning-go-web/cmd/handler"
	"github.com/facundo-vidiella/learning-go-web/internal/product"
	"github.com/gin-gonic/gin"
)

func answerPong(c *gin.Context) {
	c.String(200, "pong")
}

func main() {
	allProducts()
	var server = gin.Default()

	server.GET("/ping", answerPong)

	server.GET("/products", handler.GetProducts)
	server.GET("/products/:id", handler.GetProductById)
	server.GET("products/search/:priceGt", handler.GetProductsByPrice)
	server.POST("/products", handler.CreateProduct)

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func allProducts() {
	file, err := os.Open("./internal/products.json")
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &product.Products)
	fmt.Println(product.Products)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
