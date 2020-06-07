package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalzamora/books_server/models"
	"github.com/ricardoalzamora/books_server/services"
)

func GetAllLibros(c *gin.Context) {
	var libros []models.Libro
	err := services.GetAllLibros(&libros)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, libros)
	}
}

func CreateALibro(c *gin.Context) {
	var libro models.Libro
	c.BindJSON(&libro)
	err := services.CreateALibro(&libro)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, libro)
	}
}

func GetALibro(c *gin.Context) {
	id := c.Params.ByName("id")
	var libro models.Libro
	err := services.GetALibro(&libro, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, libro)
	}
}

func UpdateALibro(c *gin.Context) {
	var libro models.Libro
	id := c.Params.ByName("id")
	err := services.GetALibro(&libro, id)
	if err != nil {
		c.JSON(http.StatusNotFound, libro)
	}
	c.BindJSON(&libro)
	err = services.UpdateALibro(&libro, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, libro)
	}
}

func DeleteALibro(c *gin.Context) {
	var libro models.Libro
	id := c.Params.ByName("id")
	err := services.DeleteALibro(&libro, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
