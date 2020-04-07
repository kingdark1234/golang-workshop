package http

import (
	"strconv"
	"workshop-service/internals/config"
	ctrl "workshop-service/internals/controller"

	"github.com/gin-gonic/gin"
)

// ServerHTTP ...
type ServerHTTP struct {
	route       *gin.Engine
	gatewayCtrl ctrl.GatewayCtrl
	env         config.Configuration
}

// Configure ...
func (h *ServerHTTP) Configure() {
	r := h.route

	po := r.Group("/people")
	{
		// calculate quest
		po.POST("/", h.gatewayCtrl.PeopleCtrl.AddPerson)
	}

	g := r.Group("/calculate")
	{
		// calculate quest
		g.GET("/", h.gatewayCtrl.CalculateCtrl.GetResult)
		g.POST("/", h.gatewayCtrl.CalculateCtrl.Calculate)
		g.DELETE("/", h.gatewayCtrl.CalculateCtrl.ClearResult)
	}

	p := r.Group("/problem")
	{
		// calculate quest
		p.POST("/one", h.gatewayCtrl.ProblemCtrl.Multiples)
	}
}

// Start ...
func (h *ServerHTTP) Start() {
	h.Configure()

	if err := h.route.Run(":" + strconv.Itoa(h.env.Port)); err != nil {
		panic(err)
	}
}

// NewServerHTTP ...
func NewServerHTTP(gc ctrl.GatewayCtrl, env config.Configuration) *ServerHTTP {
	return &ServerHTTP{
		route:       gin.Default(),
		gatewayCtrl: gc,
		env:         env,
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
