package models

import "github.com/jinzhu/gorm"

//Pais estructura
type Pais struct {
	gorm.Model
	Codigo   string   `gorm:"not null"`
	Nombre   string   `gorm:"not null"`
	Ciudades []Ciudad `gorm:"foreignkey:CodigoPais;association_foreignkey:Codigo"`
}
