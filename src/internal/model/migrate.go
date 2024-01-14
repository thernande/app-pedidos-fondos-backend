package model

import (
	"github.com/thernande/app-pedidos-fondos-backend/internal/config/database"
	"github.com/thernande/app-pedidos-fondos-backend/internal/model/tables"
)

func Migrate() {
	db := database.New().MySQL()

	// Migrate the schema
	db.AutoMigrate(
		&tables.Usuario{},
		&tables.Asociado{},
		&tables.Fondo{},
		&tables.Restaurante{},
		&tables.Ticketera{},
		&tables.Direccion{},
		&tables.PedidoEncabezado{},
		&tables.Producto{},
		&tables.Disponibilidad{},
		&tables.Detalle{},
	)
}
