package controller

import (
	"workshop-service/internals/entity"
	"workshop-service/internals/service"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	Authors       []*entity.Author
	AuthorService *service.AuthorService
}

func (au *AuthorController) GetAuthor(c *gin.Context) {
	authors, err := au.AuthorService.ListAuthors()
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, authors)
}

func NewAuthorController() *AuthorController {
	return &AuthorController{
		Authors:       []*entity.Author{},
		AuthorService: service.NewAuthorService(),
	}
}
