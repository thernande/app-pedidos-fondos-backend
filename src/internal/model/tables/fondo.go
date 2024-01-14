package tables

import "gorm.io/gorm"

// Fondo represents the Fondo table in the database.
type Fondo struct {
	gorm.Model
	Nombre string `gorm:"unique"`
}
