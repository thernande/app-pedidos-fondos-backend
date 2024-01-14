package tables

import "gorm.io/gorm"

// Asociado represents the Asociados table in the database.
type Asociado struct {
	gorm.Model
	Documento  string
	FondoID    uint
	Ticketeras []Ticketera `gorm:"foreignKey:IdAsociado"`
}
