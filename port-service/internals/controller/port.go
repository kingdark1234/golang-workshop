package controller

import (
	"port-service/internals/entity"
	"port-service/internals/service"

	"github.com/gin-gonic/gin"
)

// PortController is ...
type PortController struct {
	Port        []*entity.Port
	PortService *service.PortService
}

// GetPorts ...
func (p *PortController) GetPorts(c *gin.Context) {
	portList, err := p.PortService.GetPort()
	if err != nil {
		panic(err)
	}
	c.JSON(200, portList)
}

// AddPorts ...
func (p *PortController) AddPorts(c *gin.Context) {
	item := &entity.Port{}
	if err := c.BindJSON(item); err != nil {
		c.AbortWithError(500, err)
		return
	}
	err := p.PortService.AddPort(item)
	if err != nil {
		panic(err)
	}
	c.JSON(200, item)
}

// NewPortController is ...
func NewPortController() *PortController {
	return &PortController{
		Port:        []*entity.Port{},
		PortService: service.NewPortService(),
	}
}
