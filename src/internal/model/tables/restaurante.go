package tables

import (
	"time"

	"gorm.io/gorm"
)

// Restaurante represents the Restaurante table in the database.
type Restaurante struct {
	gorm.Model
	Nombre            string
	Telefono          string
	Celular           string
	CorreoElectronico string
	Hora              time.Time
	Direcciones       []Direccion      `gorm:"foreignKey:IdRestaurante"`
	Productos         []Producto       `gorm:"foreignKey:IdRestaurante"`
	Disponibilidades  []Disponibilidad `gorm:"foreignKey:IdRestaurante"`
}
