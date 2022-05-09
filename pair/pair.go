package pair

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestError struct {
	Message string `json:"message"`
	Detail  string `json:"Detail"`
}

func (e *RestError) Error() string {
	return fmt.Sprintf("%s - %s", e.Message, e.Detail)
}

type Pair struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Apr       float32 `json:"apr"`
	Liquidity float32 `json:"liquidity"`
}

func (pair *Pair) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, APR: %f, Liquidity: %f", pair.ID, pair.Name, pair.Apr, pair.Liquidity)
}

func PairBind(context *gin.Context) (error, Pair) {
	var currentPair Pair
	if err := context.ShouldBindJSON(&currentPair); err != nil {
		return &RestError{"Bad Request", err.Error()}, currentPair
	}
	return nil, currentPair
}

// PostPair is the handler for the POST /pair endpoint
func PostPair(context *gin.Context) {
	// validation errors
	err, currentPair := PairBind(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	// TODO: add pair to database
	// TODO: database errors

	context.JSON(200, currentPair)
}
