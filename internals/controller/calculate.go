package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CalculateController ...
type CalculateController struct{}

var result int = 0

func (*CalculateController) Calculate(c *gin.Context) {
	paramNumberOne := c.PostForm("numberOne")
	numberOne, _ := strconv.Atoi(paramNumberOne)
	paramNumberTwo := c.PostForm("numberTwo")
	numberTwo, _ := strconv.Atoi(paramNumberTwo)
	oparator := c.PostForm("oparator")
	switch oparator {
	case "+":
		result = numberOne + numberTwo
	case "-":
		result = numberOne - numberTwo
	case "*":
		result = numberOne * numberTwo
	case "/":
		result = numberOne / numberTwo
	case "%":
		result = numberOne % numberTwo
	default:
		return
	}
	c.JSON(http.StatusOK, result)
}

func (*CalculateController) ClearResult(c *gin.Context) {
	result = 0
	c.JSON(http.StatusOK, "Success")
}

// GetPing ...
func (*CalculateController) GetResult(c *gin.Context) {
	c.JSON(http.StatusOK, result)
}

// NewCalculateController ...
func NewCalculateController() *CalculateController {
	return &CalculateController{}
}
