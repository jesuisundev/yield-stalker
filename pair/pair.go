package pair

import (
	"github.com/gin-gonic/gin"
)

func Pair(context *gin.Context) {
	context.JSON(200, "PAIR")
}
