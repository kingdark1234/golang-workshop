package controller

import (
	"workshop-service/internals/entity"
	"workshop-service/internals/service"

	"github.com/gin-gonic/gin"
)

// AuthorController .
type AuthorController struct {
	Authors       []*entity.Author
	AuthorService *service.AuthorService
}

// GetAuthor .
func (au *AuthorController) GetAuthor(c *gin.Context) {
	authors, err := au.AuthorService.ListAuthors()
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, authors)
}

// NewAuthorController .
func NewAuthorController() *AuthorController {
	return &AuthorController{
		Authors:       []*entity.Author{},
		AuthorService: service.NewAuthorService(),
	}
}
