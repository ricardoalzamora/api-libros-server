package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalzamora/books_server/models"
	"github.com/ricardoalzamora/books_server/services"
)

func GetAllGeneros(c *gin.Context) {
	var generos []models.Genero
	err := services.GetAllGeneros(&generos)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, generos)
	}
}

func CreateAGenero(c *gin.Context) {
	var genero models.Genero
	c.BindJSON(&genero)
	err := services.CreateAGenero(&genero)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, genero)
	}
}

func GetAGenero(c *gin.Context) {
	id := c.Params.ByName("id")
	var genero models.Genero
	err := services.GetAGenero(&genero, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, genero)
	}
}

func UpdateAGenero(c *gin.Context) {
	var genero models.Genero
	id := c.Params.ByName("id")
	err := services.GetAGenero(&genero, id)
	if err != nil {
		c.JSON(http.StatusNotFound, genero)
	}
	c.BindJSON(&genero)
	err = services.UpdateAGenero(&genero, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, genero)
	}
}

func DeleteAGenero(c *gin.Context) {
	var genero models.Genero
	id := c.Params.ByName("id")
	err := services.DeleteAGenero(&genero, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
