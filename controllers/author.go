package controllers

import (
	"workshop-service/models"
	"workshop-service/utils"

	"github.com/gin-gonic/gin"
)

func ListAuthor(c *gin.Context) {
	authors := []models.Author{}
	err := models.GetAllAuthor(&authors)
	if err != nil {
		utils.RespondJSON(c, 404, authors)
	} else {
		utils.RespondJSON(c, 200, authors)
	}
}

func AddNewAuthor(c *gin.Context) {
	author := models.Author{}
	c.BindJSON(&author)
	err := models.AddNewAuthor(&author)
	if err != nil {
		utils.RespondJSON(c, 404, author)
	} else {
		utils.RespondJSON(c, 200, author)
	}
}
