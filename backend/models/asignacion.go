package models

import "time"

type Asignacion struct {
	ID          uint      `gorm:"primaryKey"`
	OperadorID  string    `gorm:"not null"`
	SolicitudID uint      `gorm:"not null"`
	Fecha       time.Time `gorm:"autoCreateTime"`
	Estado      string    `gorm:"size:20;not null;default:'pendiente'"` // pendiente, aprobado, rechazado, mora, al_dia
}
