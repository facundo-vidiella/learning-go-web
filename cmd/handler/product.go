package handler

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/facundo-vidiella/learning-go-web/internal/domain"
	"github.com/facundo-vidiella/learning-go-web/internal/product"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(200, product.Products)
}

func GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	found := false
	for _, n := range product.Products {
		if n.ID == id {
			c.JSON(200, n)
			found = true
		}
	}
	if !found {
		c.JSON(400, `errorMsg: No product found with that id`)
	}
}

func GetProductsByPrice(c *gin.Context) {

	priceGt, err := strconv.ParseFloat(c.Param("priceGt"), 64)
	if err != nil {
		log.Fatal(err)
	}
	var validProducts []domain.Product

	for _, n := range product.Products {
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
	if splitedCurrentTime[0] > splitedDate[2] {
		return false
	}
	return re.MatchString(date)
}

func CreateProduct(c *gin.Context) {
	var req domain.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.ID == 0 {
		req.ID = len(product.Products) + 1
	}

	for _, n := range product.Products {
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

	product.Products = append(product.Products, req)

	c.JSON(200, gin.H{
		"success": "product succesfully created",
	})

}
