package tables

import (
	"time"

	"gorm.io/gorm"
)

// PedidoEncabezado represents the Pedido_Encabezado table in the database.
type PedidoEncabezado struct {
	gorm.Model
	IdUsuario   uint
	IdTicketera uint
	IdDireccion uint
	ValorTotal  float64
	MedioDePago string
	Fecha       time.Time
	Hora        time.Time
	Detalles    []Detalle `gorm:"foreignKey:IdEncabezado"`
}
