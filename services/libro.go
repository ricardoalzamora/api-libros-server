package services

import (
	"fmt"

	"github.com/ricardoalzamora/books_server/config"
	"github.com/ricardoalzamora/books_server/models"
)

func GetAllLibros(libros *[]models.Libro) (err error) {
	if err = config.DB.Find(libros).Error; err != nil {
		return err
	}
	return nil
}

func CreateALibro(libro *models.Libro) (err error) {
	if err = config.DB.Create(libro).Error; err != nil {
		return err
	}
	return nil
}

func GetALibro(libro *models.Libro, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(libro).Error; err != nil {
		return err
	}
	return nil
}

func UpdateALibro(libro *models.Libro, id string) (err error) {
	fmt.Println(libro)
	config.DB.Save(libro)
	return nil
}

func DeleteALibro(libro *models.Libro, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(libro)
	return nil
}
