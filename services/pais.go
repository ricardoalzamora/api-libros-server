package services

import (
	"fmt"

	"github.com/ricardoalzamora/books_server/config"
	"github.com/ricardoalzamora/books_server/models"
)

func GetAllPaises(paises *[]models.Pais) (err error) {
	if err = config.DB.Find(paises).Error; err != nil {
		return err
	}
	return nil
}

func CreateAPais(pais *models.Pais) (err error) {
	if err = config.DB.Create(pais).Error; err != nil {
		return err
	}
	return nil
}

func GetAPais(pais *models.Pais, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(pais).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAPais(pais *models.Pais, id string) (err error) {
	fmt.Println(pais)
	config.DB.Save(pais)
	return nil
}

func DeleteAPais(pais *models.Pais, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(pais)
	return nil
}
