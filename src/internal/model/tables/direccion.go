package tables

import (
	"gorm.io/gorm"
)

// Direccion represents the Direcciones table in the database.
type Direccion struct {
	gorm.Model
	IdUsuario     uint
	IdRestaurante uint
	Departamento  string
	Ciudad        string
	Barrio        string
	Complemento1  string
	Complemento2  string
	Direccion     string
}
