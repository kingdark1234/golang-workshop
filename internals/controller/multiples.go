package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProblemController ...
type ProblemController struct{}

func isIntegral(val float32) bool {
	return val == float32(int(val))
}

// Multiples ...
func (*ProblemController) Multiples(c *gin.Context) {
	param := c.PostForm("size")
	size, _ := strconv.Atoi(param)
	var arrayResult int
	for i := 1; i < size; i++ {
		var multiWIthThree, multiWIthFive float32
		multiWIthThree = float32(i) / 3
		multiWIthFive = float32(i) / 5
		if isIntegral(multiWIthThree) || isIntegral(multiWIthFive) {
			arrayResult = arrayResult + i
		}
	}
	c.JSON(http.StatusOK, arrayResult)
}

// NewProblemController ...
func NewProblemController() *ProblemController {
	return &ProblemController{}
}
