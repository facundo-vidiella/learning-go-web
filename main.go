package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id" omitempty`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func answerPong(c *gin.Context) {
	c.String(200, "pong")
}

func getProducts(c *gin.Context) {
	c.JSON(200, products)
}

func getProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	found := false
	for _, n := range products {
		if n.ID == id {
			c.JSON(200, n)
			found = true
		}
	}
	if !found {
		c.JSON(400, `errorMsg: No product found with that id`)
	}
}

func getProductsByPrice(c *gin.Context) {

	priceGt, err := strconv.ParseFloat(c.Param("priceGt"), 64)
	if err != nil {
		log.Fatal(err)
	}
	var validProducts []Product

	for _, n := range products {
		if n.Price > priceGt {
			validProducts = append(validProducts, n)
		}
	}

	if len(validProducts) == 0 {
		c.JSON(400, `errorMsg: no products found with a price bigger than the one specified`)
		return
	}
	jsonResponse := map[string]interface{}{
		"total":    len(validProducts),
		"products": validProducts,
	}
	c.JSON(200, jsonResponse)
}

func validateDate(date string) bool {
	re := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
	currentTime := time.Now().Format("2006-01-02")
	splitedCurrentTime := strings.Split(currentTime, "-")
	splitedDate := strings.Split(date, "/")
	fmt.Println(splitedCurrentTime)
	if splitedCurrentTime[0] > splitedDate[2] {
		return false
	}
	return re.MatchString(date)
}

func createProduct(c *gin.Context) {
	var req Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.ID == 0 {
		req.ID = len(products) + 1
	}

	for _, n := range products {
		if n.CodeValue == req.CodeValue {
			c.JSON(404, gin.H{
				"error": "There is another product with the same code_value",
			})
			return
		}
	}

	if !validateDate(req.Expiration) {
		c.JSON(400, gin.H{
			"error": "Invalid date format or input",
		})
		return
	}

	products = append(products, req)

	c.JSON(200, gin.H{
		"success": "product succesfully created",
	})

}

func main() {
	allProducts()
	var server = gin.Default()

	server.GET("/ping", answerPong)

	server.GET("/products", getProducts)
	server.GET("/products/:id", getProductById)
	server.GET("products/search/:priceGt", getProductsByPrice)
	server.POST("/products", createProduct)

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func allProducts() {
	file, err := os.Open("products.json")
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &products)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
