package models

import "github.com/jinzhu/gorm"

//Ciudad estructura
type Ciudad struct {
	gorm.Model
	Codigo          string `gorm:"not null"`
	Nombre          string `gorm:"not null"`
	CodigoPaisRefer string `gorm:"not null"`
	Pais            Pais
}
