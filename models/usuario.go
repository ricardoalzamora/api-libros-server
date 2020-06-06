package models

import (
	"github.com/jinzhu/gorm"
)

//Usuario estructura
type Usuario struct {
	gorm.Model
	Nombre             string `gorm:"not null"`
	Apellido           string `gorm:"not null"`
	GeneroID           uint   `gorm:"not null"`
	Genero             Genero
	CiudadNacimientoID uint `gorm:"not null"`
	CiudadNacimiento   Ciudad
	LibrosDeAutoria    []*Libro `gorm:"many2many:usuario_libros"`
}
