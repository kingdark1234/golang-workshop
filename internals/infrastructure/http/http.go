package http

import (
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

	g := r.Group("/author")
	{
		g.GET("/getAll", h.gatewayCtrl.AuthorCtrl.GetAuthor)
	}
}

// Start ...
func (h *Server) Start() {
	h.Configure()

	if err := h.route.Run(":3000"); err != nil {
		panic(err)
	}
}

// NewHTTPServer ...
func NewHTTPServer(g ctrl.GatewayCtrl, env config.Configuration) *Server {
	return &Server{
		route: gin.Default(),
		// gatewayCtrl: g,
		// env:         env,
	}
}
