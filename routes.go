package main

func Routes() {

	ProductsGroup.GET("/products", getProducts)
	ProductsGroup.GET("/products/:id", getProductById)
	ProductsGroup.GET("products/search/:priceGt", getProductsByPrice)
	ProductsGroup.POST("/products", createProduct)

}
