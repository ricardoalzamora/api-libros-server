package models

import "github.com/jinzhu/gorm"

//Editorial estructura
type Editorial struct {
	gorm.Model
	Nombre       string `gorm:"not null"`
	Codigo       string `gorm:"not null"`
	CodigoCiudad string `gorm:"not null"`
	Ciudad       Ciudad
}
