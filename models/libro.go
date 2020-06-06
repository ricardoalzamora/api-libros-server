package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Libro estructura
type Libro struct {
	gorm.Model
	Codigo           string
	Nombre           string
	FechaPublicacion time.Time  `gorm:"not null"`
	Autores          []*Usuario `gorm:"many2many:usuario_libros"`
	CasaEditorialID  uint       `gorm:"not null"`
	CasaEditorial    Editorial
}
