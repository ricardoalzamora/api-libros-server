package services

import (
	"fmt"

	"github.com/ricardoalzamora/books_server/config"
	"github.com/ricardoalzamora/books_server/models"
)

func GetAllGeneros(generos *[]models.Genero) (err error) {
	if err = config.DB.Find(generos).Error; err != nil {
		return err
	}
	return nil
}

func CreateAGenero(genero *models.Genero) (err error) {
	if err = config.DB.Create(genero).Error; err != nil {
		return err
	}
	return nil
}

func GetAGenero(genero *models.Genero, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(genero).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAGenero(genero *models.Genero, id string) (err error) {
	fmt.Println(genero)
	config.DB.Save(genero)
	return nil
}

func DeleteAGenero(genero *models.Genero, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(genero)
	return nil
}
