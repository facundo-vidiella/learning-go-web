package main

import "github.com/gin-gonic/gin"

func route(){
	
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.Data(200, "text/html; charset=utf-8", []byte("pong"))
	})

	router.Run()
}
