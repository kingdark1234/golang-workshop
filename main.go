package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type person struct {
	id        int
	firstName string
	surName   string
	age       int
}

type persons []person

func print(c *gin.Context) {
	var p *person
	c.BindJSON(&p)
	fmt.Println(c)
	c.JSON(http.StatusOK, gin.H{"firstName": p.firstName})
}

var result int = 0

func getResult(c *gin.Context) {
	c.JSON(http.StatusOK, result)
}

func calculate(c *gin.Context) {
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
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}

func clearResult(c *gin.Context) {
	result = 0
	c.JSON(http.StatusOK, "Success")
}

func isIntegral(val float32) bool {
	return val == float32(int(val))
}

func multiples(c *gin.Context) {
	param := c.PostForm("size")
	size, _ := strconv.Atoi(param)
	var arrayResult int = 0
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

func main() {
	route := gin.Default()

	// people quest
	route.GET("/peoples/", func(c *gin.Context) {

	})
	route.GET("/people/:id", func(c *gin.Context) {

	})

	route.POST("/people/", print)

	route.PUT("/people/", func(c *gin.Context) {

	})
	route.DELETE("/people/", func(c *gin.Context) {

	})

	// calculate quest
	route.GET("/calculate/", getResult)
	route.POST("/calculate/", calculate)
	route.DELETE("/calculate/", clearResult)
	// problem1 Multiples of 3 and 5
	route.POST("/problem/One", multiples)

	// problem2 Even Fibonacci numbers
	route.GET("/problem/Two", func(c *gin.Context) {

	})

	// problem3 Largest prime factor
	route.GET("/problem/Three", func(c *gin.Context) {

	})

	// problem4 Largest palindrome product
	route.GET("/problem/Four", func(c *gin.Context) {

	})

	// problem5 Smallest multiple
	route.GET("/problem/Five", func(c *gin.Context) {

	})

	route.Run()
}
