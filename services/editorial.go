package services

import (
	"fmt"

	"github.com/ricardoalzamora/books_server/config"
	"github.com/ricardoalzamora/books_server/models"
)

func GetAllEditoriales(editoriales *[]models.Editorial) (err error) {
	if err = config.DB.Find(editoriales).Error; err != nil {
		return err
	}
	return nil
}

func CreateAEditorial(editorial *models.Editorial) (err error) {
	if err = config.DB.Create(editorial).Error; err != nil {
		return err
	}
	return nil
}

func GetAEditorial(editorial *models.Editorial, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(editorial).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAEditorial(editorial *models.Editorial, id string) (err error) {
	fmt.Println(editorial)
	config.DB.Save(editorial)
	return nil
}

func DeleteAEditorial(editorial *models.Editorial, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(editorial)
	return nil
}
