package controller

import (
	"github.com/gin-gonic/gin"
)

// PingController ...
type PingController struct{}

// GetPing ...
func (*PingController) GetPing(c *gin.Context) {
	c.String(200, "Pong")
}

// NewPingController ...
func NewPingController() *PingController {
	return &PingController{}
}
