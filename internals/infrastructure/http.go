package http

import (
	ctrl "workshop-service/internals/controller"

	"github.com/gin-gonic/gin"
)

// ServerHTTP ...
type ServerHTTP struct {
	route         *gin.Engine
	calculateCtrl *ctrl.CalculateController
}

// Configure ...
func (h *ServerHTTP) Configure() {
	r := h.route
	g := r.Group("/calculate")
	{
		// calculate quest
		g.GET("/calculate/", h.calculateCtrl.GetResult)
		g.POST("/calculate/", h.calculateCtrl.Calculate)
		g.DELETE("/calculate/", h.calculateCtrl.ClearResult)
	}
}

// Start ...
func (h *ServerHTTP) Start() {
	h.Configure()
	if err := h.route.Run(":3000"); err != nil {
		panic(err)
	}
}

// NewServerHTTP ...
func NewServerHTTP() *ServerHTTP {
	return &ServerHTTP{
		route:         gin.Default(),
		calculateCtrl: ctrl.NewPingController(),
	}
}

// func main() {
// 	route := gin.Default()

// 	// people quest
// 	route.GET("/peoples/", func(c *gin.Context) {

// 	})
// 	route.GET("/people/:id", func(c *gin.Context) {

// 	})

// 	route.POST("/people/", print)

// 	route.PUT("/people/", func(c *gin.Context) {

// 	})
// 	route.DELETE("/people/", func(c *gin.Context) {

// 	})

// 	// problem1 Multiples of 3 and 5
// 	route.POST("/problem/One", multiples)

// 	// problem2 Even Fibonacci numbers
// 	route.GET("/problem/Two", func(c *gin.Context) {

// 	})

// 	// problem3 Largest prime factor
// 	route.GET("/problem/Three", func(c *gin.Context) {

// 	})

// 	// problem4 Largest palindrome product
// 	route.GET("/problem/Four", func(c *gin.Context) {

// 	})

// 	// problem5 Smallest multiple
// 	route.GET("/problem/Five", func(c *gin.Context) {

// 	})

// 	route.Run()
// }
