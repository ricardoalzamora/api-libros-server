package services

import (
	"fmt"

	"github.com/ricardoalzamora/books_server/config"
	"github.com/ricardoalzamora/books_server/models"
)

func GetAllUsuarios(usuarios *[]models.Usuario) (err error) {
	if err = config.DB.Find(usuarios).Error; err != nil {
		return err
	}
	return nil
}

func CreateAUsuario(usuario *models.Usuario) (err error) {
	if err = config.DB.Create(usuario).Error; err != nil {
		return err
	}
	return nil
}

func GetAUsuario(usuario *models.Usuario, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(usuario).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAUsuario(usuario *models.Usuario, id string) (err error) {
	fmt.Println(usuario)
	config.DB.Save(usuario)
	return nil
}

func DeleteAUsuario(usuario *models.Usuario, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(usuario)
	return nil
}
