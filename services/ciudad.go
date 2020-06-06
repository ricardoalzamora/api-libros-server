package services

import (
	"fmt"

	"github.com/ricardoalzamora/books_server/config"
	"github.com/ricardoalzamora/books_server/models"
)

func GetAllCiudades(ciudades *[]models.Ciudad) (err error) {
	if err = config.DB.Find(ciudades).Error; err != nil {
		return err
	}
	return nil
}

func CreateACiudad(ciudad *models.Ciudad) (err error) {
	if err = config.DB.Create(ciudad).Error; err != nil {
		return err
	}
	return nil
}

func GetACiudad(ciudad *models.Ciudad, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(ciudad).Error; err != nil {
		return err
	}
	return nil
}

func UpdateACiudad(ciudad *models.Ciudad, id string) (err error) {
	fmt.Println(ciudad)
	config.DB.Save(ciudad)
	return nil
}

func DeleteACiudad(ciudad *models.Ciudad, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(ciudad)
	return nil
}
