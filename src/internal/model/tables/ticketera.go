package tables

import (
	"time"

	"gorm.io/gorm"
)

// Ticketera represents the Ticketeras table in the database.
type Ticketera struct {
	gorm.Model
	IdAsociado uint
	Cantidad   int
	Vigencia   time.Time
	Valor      float64
	Pedidos    []PedidoEncabezado `gorm:"foreignKey:IdTicketera"`
}
