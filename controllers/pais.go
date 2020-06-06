package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalzamora/books_server/models"
	"github.com/ricardoalzamora/books_server/services"
)

func GetPaises(c *gin.Context) {
	var paises []models.Pais
	err := services.GetAllPaises(&paises)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, paises)
	}
}

func CreateAPais(c *gin.Context) {
	var pais models.Pais
	c.BindJSON(&pais)
	err := services.CreateAPais(&pais)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pais)
	}
}

func GetAPais(c *gin.Context) {
	id := c.Params.ByName("id")
	var pais models.Pais
	err := services.GetAPais(&pais, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pais)
	}
}

func UpdateAPais(c *gin.Context) {
	var pais models.Pais
	id := c.Params.ByName("id")
	err := services.GetAPais(&pais, id)
	if err != nil {
		c.JSON(http.StatusNotFound, pais)
	}
	c.BindJSON(&pais)
	err = services.UpdateAPais(&pais, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pais)
	}
}

func DeleteAPais(c *gin.Context) {
	var pais models.Pais
	id := c.Params.ByName("id")
	err := services.DeleteAPais(&pais, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
