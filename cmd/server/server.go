package main

import (
	"fmt"
	"strconv"

	"workshop-service/internals/model"
	"workshop-service/internals/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	arrPerson := []model.Person{}

	route := gin.Default()

	// people quest
	route.GET("/people/all", func(c *gin.Context) {
		resp := arrPerson
		c.JSON(200, resp)
	})

	route.GET("/people/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := &model.Person{}
		for _, v := range arrPerson {
			if strconv.Itoa(v.ID) == id {
				fmt.Println("found")
				*user = v
				c.JSON(200, *user)
				break
			}
		}
	})

	route.POST("/people/", func(c *gin.Context) {
		route := &model.Person{}
		c.BindJSON(route)
		arrPerson = append(arrPerson, *route)
		// controller.CreatePersons(route, arrPerson)
		fmt.Println(arrPerson)

		c.JSON(200, route)
	})

	route.PUT("/people/:id", func(c *gin.Context) {
		id := c.Param("id")
		idNum, _ := strconv.Atoi(id)
		route := &model.Person{
			ID: idNum,
		}
		c.BindJSON(route)
		for k, v := range arrPerson {
			if strconv.Itoa(v.ID) == id {
				arrPerson[k] = *route
				break
			}
		}
		c.JSON(200, arrPerson)
	})

	route.DELETE("/people/:id", func(c *gin.Context) {
		id := c.Param("id")
		idNum, _ := strconv.Atoi(id)
		route := &model.Person{
			ID: idNum,
		}
		c.BindJSON(route)
		for k, v := range arrPerson {
			if strconv.Itoa(v.ID) == id {
				arrPerson[k] = arrPerson[len(arrPerson)-1]
				arrPerson[len(arrPerson)-1] = model.Person{}
				arrPerson = arrPerson[:len(arrPerson)-1]
				break
			}
		}
		c.JSON(200, arrPerson)
	})

	// ----------------------------------------------------

	var currentResult float32
	// calculate quest
	route.GET("/calculate/", func(c *gin.Context) {
		// TODO: return current result
		c.JSON(200, currentResult)
	})
	route.POST("/calculate/", func(c *gin.Context) {
		// TODO: return result input:{number1:1,operation:'+',number2:2}
		// support + - * / %
		// response = 3
		route := &model.Calculate{}
		c.BindJSON(route)

		firstNumber := route.FirstNumber
		secondNumber := route.SecondNumber

		if route.Operand == "+" {
			// utils.Sum(firstNumber, secondNumber, currentResult)
			// c.JSON(200, currentResult)
			result := utils.Sum(firstNumber, secondNumber)
			currentResult = result
			c.JSON(200, result)
		}
		if route.Operand == "-" {
			result := utils.Difference(firstNumber, secondNumber)
			currentResult = result
			c.JSON(200, result)
		}
		if route.Operand == "*" {
			result := utils.Product(firstNumber, secondNumber)
			currentResult = result
			c.JSON(200, result)
		}
		if route.Operand == "/" {
			result := utils.Quotient(firstNumber, secondNumber)
			currentResult = result
			c.JSON(200, result)
		}
		if route.Operand == "%" {
			result := utils.Remainder(firstNumber, secondNumber)
			currentResult = result
			c.JSON(200, result)
		}
	})
	route.DELETE("/calculate/", func(c *gin.Context) {
		// TODO: clear Result
		currentResult = utils.ClearValue()
		c.JSON(200, currentResult)
	})

	// problem1 Multiples of 3 and 5
	route.POST("/problem/One", func(c *gin.Context) {

	})

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

	err := route.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("server start...")

}
