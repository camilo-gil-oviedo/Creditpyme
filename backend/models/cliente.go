package models

import "time"

type Cliente struct {
	ID            uint      `gorm:"primaryKey"`
	Nombre        string    `gorm:"size:100;not null"`
	Apellido      string    `gorm:"size:100;not null"`
	Correo        string    `gorm:"size:100;unique;not null"`
	Contrasena    string    `gorm:"size:255;not null"`
	FechaRegistro time.Time `gorm:"autoCreateTime"`
}

type Empresa struct {
	ID        uint   `gorm:"primaryKey"`
	ClienteID uint   `gorm:"uniqueIndex"`       // 1:1 con Cliente
	Nombre    string `gorm:"size:100;not null"` // Nombre de la empresa
	Direccion string `gorm:"size:255;not null"` // Dirección
	Ciudad    string `gorm:"size:100;not null"` // Ciudad
}

type SolicitudCredito struct {
	ID              uint    `gorm:"primaryKey"`
	ClienteID       uint    `gorm:"not null"`          // FK a Cliente
	MontoSolicitado float64 `gorm:"not null"`          // Monto del préstamo
	PlazoMeses      int     `gorm:"not null"`          // Plazo en meses
	DestinoCredito  string  `gorm:"size:255;not null"` // Destino del dinero
}
