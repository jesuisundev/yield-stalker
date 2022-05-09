package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jesuisundev/yield-stalker/pair"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(context *gin.Context) { hello(context) })

	router.POST("/pair", func(context *gin.Context) { pair.PostPair(context) })

	router.Run()
}
