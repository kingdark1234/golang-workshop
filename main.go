package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	// people quest
	route.GET("/people/all", func(c *gin.Context) {

	})
	route.GET("/people/:id", func(c *gin.Context) {

	})
	route.POST("/people/", func(c *gin.Context) {

	})
	route.PUT("/people/", func(c *gin.Context) {

	})
	route.DELETE("/people/", func(c *gin.Context) {

	})

	// calculate quest
	route.GET("/calculate/", func(c *gin.Context) {
		// TODO: return current result
	})
	route.POST("/calculate/", func(c *gin.Context) {
		// TODO: return result
	})
	route.DELETE("/calculate/", func(c *gin.Context) {
		// TODO: clear Result
	})

	// calculate quest
	route.GET("/calculate/", func(c *gin.Context) {
		// TODO: return current result
	})
	route.POST("/calculate/", func(c *gin.Context) {
		// TODO: return result input:{number1:1,oparation:'+',number2:2}
		// support + - * / %
		// response = 3
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

	route.Run()
}
