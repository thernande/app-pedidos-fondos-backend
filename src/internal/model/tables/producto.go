package tables

import "gorm.io/gorm"

// Producto represents the Producto table in the database.
type Producto struct {
	gorm.Model
	IdRestaurante    uint
	IdDisponibilidad uint
	Nombre           string
	Categoria        string
	Descripcion      string
	Descuento        float64
	Valor            float64
	Detalles         []Detalle `gorm:"foreignKey:IdProducto"`
}
