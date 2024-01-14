package tables

import (
	"gorm.io/gorm"
)

// Detalle represents the Detalle table in the database.
type Detalle struct {
	gorm.Model
	IdUsuario     uint
	IdProducto    uint
	IdEncabezado  uint
	Cantidad      uint
	Observaciones string
}
