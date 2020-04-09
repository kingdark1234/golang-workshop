package http

import (
	"strconv"
	"workshop-service/internals/config"
	ctrl "workshop-service/internals/controller"

	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	route       *gin.Engine
	gatewayCtrl ctrl.GatewayCtrl
	env         config.Configuration
}

// Configure ...
func (h *Server) Configure() {
	r := h.route
	r.GET("/ping", h.gatewayCtrl.PingCtrl.GetPing)

	g := r.Group("/product")
	{
		g.GET("/getAll", h.gatewayCtrl.ProductCtrl.GetProducts)
		g.POST("/add", h.gatewayCtrl.ProductCtrl.AddProduct)
		g.PUT("/update/:id", h.gatewayCtrl.ProductCtrl.UpdateProduct)
	}
}

// Start ...
func (h *Server) Start() {
	h.Configure()

	if err := h.route.Run(":" + strconv.Itoa(h.env.Port)); err != nil {
		panic(err)
	}
}

// NewHTTPServer ...
func NewHTTPServer(g ctrl.GatewayCtrl, env config.Configuration) *Server {
	return &Server{
		route:       gin.Default(),
		gatewayCtrl: g,
		env:         env,
	}
}
