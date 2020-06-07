package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalzamora/books_server/models"
	"github.com/ricardoalzamora/books_server/services"
)

func CreateAUsuario(c *gin.Context) {
	var usuario models.Usuario
	c.BindJSON(&usuario)
	err := services.CreateAUsuario(&usuario)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, usuario)
	}
}

func GetAUsuario(c *gin.Context) {
	id := c.Params.ByName("id")
	var usuario models.Usuario
	err := services.GetAUsuario(&usuario, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, usuario)
	}
}

func UpdateAUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	err := services.GetAUsuario(&usuario, id)
	if err != nil {
		c.JSON(http.StatusNotFound, usuario)
	}
	c.BindJSON(&usuario)
	err = services.UpdateAUsuario(&usuario, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, usuario)
	}
}

func DeleteAUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	err := services.DeleteAUsuario(&usuario, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
