package http

import (
	ctrl "port-service/internals/controller"

	"github.com/gin-gonic/gin"
)

// ServerHTTP ...
type ServerHTTP struct {
	route    *gin.Engine
	pingCtrl *ctrl.PingController
	portCtrl *ctrl.PortController
}

// Configure ...
func (h *ServerHTTP) Configure() {
	r := h.route
	r.GET("/ping", h.pingCtrl.GetPing)

	g := r.Group("/port")
	{
		g.GET("/getAll", h.portCtrl.GetPorts)
		g.POST("/add", h.portCtrl.AddPorts)
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
		route:    gin.Default(),
		pingCtrl: ctrl.NewPingController(),
		portCtrl: ctrl.NewPortController(),
	}
}
