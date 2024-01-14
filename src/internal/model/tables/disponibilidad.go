package tables

import (
	"time"

	"gorm.io/gorm"
)

// Disponibilidad represents the Disponibilidad table in the database.
type Disponibilidad struct {
	gorm.Model
	IdRestaurante uint
	Nombre        string
	HoraInicio    time.Time
	HoraFinal     time.Time
}
