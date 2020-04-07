package controller

import (
	"workshop-service/internals/entity"
	repo "workshop-service/internals/repository"

	"github.com/gin-gonic/gin"
)

// PeopleController ...
type PeopleController struct {
	Peoples     []*entity.People
	gatewayRepo repo.GatewayRepo
}

// AddPerson ...
func (p *PeopleController) AddPerson(c *gin.Context) {
	item := &entity.AddPeople{}
	if err := c.BindJSON(item); err != nil {
		c.AbortWithError(500, err)
		return
	}
	err := p.gatewayRepo.PeopleRepo.Add(item)
	if err != nil {
		panic(err)
	}
	c.JSON(200, item)
}

// NewPeopleController ...
func NewPeopleController() *PeopleController {
	return &PeopleController{}
}
