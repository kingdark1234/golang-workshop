package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CalculateController ...
type CalculateController struct{}

var result int = 0

// Calculate ...
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

// ClearResult ...
func (*CalculateController) ClearResult(c *gin.Context) {
	result = 0
	c.JSON(http.StatusOK, "Success")
}

// GetResult ...
func (*CalculateController) GetResult(c *gin.Context) {
	c.JSON(http.StatusOK, result)
}

// NewCalculateController ...
func NewCalculateController() *CalculateController {
	return &CalculateController{}
}
