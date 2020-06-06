package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ricardoalzamora/books_server/config"
	"github.com/ricardoalzamora/books_server/models"
	"github.com/ricardoalzamora/books_server/routes"
)

var err error

func main() {

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	config.DB.SingularTable(true)
	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Ciudad{}, &models.Editorial{}, &models.Genero{}, &models.Libro{}, &models.Pais{}, &models.Usuario{})

	r := routes.SetupRouter()
	// running
	r.Run()
}
