package tables

import (
	"gorm.io/gorm"
)

// Usuario represents the Usuarios table in the database.
type Usuario struct {
	gorm.Model
	Documento   string `gorm:"unique"`
	Nombres     string
	Apellidos   string
	Telefono    string
	Correo      string `gorm:"unique"`
	Empresa     string
	Direcciones []Direccion        `gorm:"foreignKey:IdUsuario"`
	PedidoEnc   []PedidoEncabezado `gorm:"foreignKey:IdUsuario"`
}
