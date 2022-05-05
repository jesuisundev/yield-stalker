package main

import (
	"github.com/gin-gonic/gin"
)

func hello(context *gin.Context) {
	context.JSON(200, "Hello World")
}
