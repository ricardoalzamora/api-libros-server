package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalzamora/books_server/models"
	"github.com/ricardoalzamora/books_server/services"
)

func GetAllCiudades(c *gin.Context) {
	var ciudades []models.Ciudad
	err := services.GetAllCiudades(&ciudades)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ciudades)
	}
}

func CreateACiudad(c *gin.Context) {
	var ciudad models.Ciudad
	c.BindJSON(&ciudad)
	err := services.CreateACiudad(&ciudad)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ciudad)
	}
}

func GetACiudad(c *gin.Context) {
	id := c.Params.ByName("id")
	var ciudad models.Ciudad
	err := services.GetACiudad(&ciudad, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ciudad)
	}
}

func UpdateACiudad(c *gin.Context) {
	var ciudad models.Ciudad
	id := c.Params.ByName("id")
	err := services.GetACiudad(&ciudad, id)
	if err != nil {
		c.JSON(http.StatusNotFound, ciudad)
	}
	c.BindJSON(&ciudad)
	err = services.UpdateACiudad(&ciudad, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ciudad)
	}
}

func DeleteACiudad(c *gin.Context) {
	var ciudad models.Ciudad
	id := c.Params.ByName("id")
	err := services.DeleteACiudad(&ciudad, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
