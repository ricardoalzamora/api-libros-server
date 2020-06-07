package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalzamora/books_server/models"
	"github.com/ricardoalzamora/books_server/services"
)

func GetAllEditoriales(c *gin.Context) {
	var editoriales []models.Editorial
	err := services.GetAllEditoriales(&editoriales)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, editoriales)
	}
}

func CreateAEditorial(c *gin.Context) {
	var editorial models.Editorial
	c.BindJSON(&editorial)
	err := services.CreateAEditorial(&editorial)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, editorial)
	}
}

func GetAEditorial(c *gin.Context) {
	id := c.Params.ByName("id")
	var editorial models.Editorial
	err := services.GetAEditorial(&editorial, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, editorial)
	}
}

func UpdateAEditorial(c *gin.Context) {
	var editorial models.Editorial
	id := c.Params.ByName("id")
	err := services.GetAEditorial(&editorial, id)
	if err != nil {
		c.JSON(http.StatusNotFound, editorial)
	}
	c.BindJSON(&editorial)
	err = services.UpdateAEditorial(&editorial, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, editorial)
	}
}

func DeleteAEditorial(c *gin.Context) {
	var editorial models.Editorial
	id := c.Params.ByName("id")
	err := services.DeleteAEditorial(&editorial, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
