package handlers

import (
	"net/http"
	"time"

	"github.com/Camilo/creditPYMESbackend/db"
	"github.com/Camilo/creditPYMESbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegistrarClienteCompleto maneja cliente + empresa + solicitud
func RegistrarClienteCompleto(c *gin.Context) {
	var payload struct {
		Cliente   models.Cliente          `json:"cliente"`
		Empresa   models.Empresa          `json:"empresa"`
		Solicitud models.SolicitudCredito `json:"solicitud"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos", "detalle": err.Error()})
		return
	}

	// Establecer fecha de registro del cliente
	payload.Cliente.FechaRegistro = time.Now()

	// Conexión a la base de datos
	database, err := db.Connect("host=localhost user=postgres password='Jjosee123&' dbname=fintech port=5432 sslmode=disable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo conectar a la base de datos", "detalle": err.Error()})
		return
	}

	// Guardamos todo en una transacción para asegurar consistencia
	err = database.Transaction(func(tx *gorm.DB) error {
		// Guardar cliente
		if res := tx.Create(&payload.Cliente); res.Error != nil {
			return res.Error
		}

		// Asociar empresa al cliente y guardar
		payload.Empresa.ClienteID = payload.Cliente.ID
		if res := tx.Create(&payload.Empresa); res.Error != nil {
			return res.Error
		}

		// Asociar solicitud al cliente y guardar
		payload.Solicitud.ClienteID = payload.Cliente.ID
		if res := tx.Create(&payload.Solicitud); res.Error != nil {
			return res.Error
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo registrar toda la información", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "registro exitoso",
		"cliente":   payload.Cliente,
		"empresa":   payload.Empresa,
		"solicitud": payload.Solicitud,
	})
}
