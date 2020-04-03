package people

import "github.com/gin-gonic/gin"

import "fmt"

type Person struct {
	id int
	firstName string
	surName string
	age  int
}

type persons []struct

// addperson ...
func AddPerson (c *gin.Context) {
	fmt.Println("Hello World")
}