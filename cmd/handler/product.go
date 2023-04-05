package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/facundo-vidiella/learning-go-web/internal/domain"
	"github.com/facundo-vidiella/learning-go-web/internal/product"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	repo := new(product.Repository)
	products := repo.GetProducts()
	c.JSON(200, products)
}

func GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	repo := new(product.Repository)

	product, err := repo.GetProductById(id)

	if err != nil {
		c.JSON(400, `errorMsg: No product found with that id`)
	}
	c.JSON(200, product)
}

func GetProductsByPrice(c *gin.Context) {

	priceGt, err := strconv.ParseFloat(c.Param("priceGt"), 64)
	if err != nil {
		log.Fatal(err)
	}
	repo := new(product.Repository)

	validProducts := repo.GetProductsByPrice(priceGt)

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

func CreateProduct(c *gin.Context) {
	var req domain.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := new(product.Repository)

	message, err := repo.CreateProduct(req)

	if err != nil {
		errMessage := fmt.Sprintf("%s: %s", err, message)
		c.JSON(400, errMessage)
	}

	c.JSON(200, gin.H{
		"success": message,
	})

}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	var req domain.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := new(product.Repository)
	message, err := repo.UpdateProduct(id, req)
	if err != nil {
		errMessage := fmt.Sprintf("%s: %s", err, message)

		c.JSON(400, errMessage)
	}
	c.JSON(200, message)
}

func UpdateProductField(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	var req domain.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := new(product.Repository)
	message, err := repo.UpdateProductField(id, req)
	if err != nil {
		errMessage := fmt.Sprintf("%s: %s", err, message)

		c.JSON(400, errMessage)
	}
	c.JSON(200, message)
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	repo := new(product.Repository)

	err = repo.DeleteProduct(id)

	if err != nil {
		c.JSON(400, gin.H{
			"Error": err,
		})
	}

	c.JSON(204, "")
}
