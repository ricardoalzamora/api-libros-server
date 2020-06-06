package models

import "github.com/jinzhu/gorm"

//Genero structura
type Genero struct {
	gorm.Model
	Nombre string `gorm:"not null"`
}
