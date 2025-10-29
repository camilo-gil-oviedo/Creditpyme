package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Camilo/creditPYMESbackend/auth"
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

	payload.Cliente.FechaRegistro = time.Now()

	database, err := db.Connect("host=localhost user=postgres password='Jjosee123&' dbname=fintech port=5433 sslmode=disable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo conectar a la base de datos", "detalle": err.Error()})
		return
	}

	err = database.Transaction(func(tx *gorm.DB) error {
		if res := tx.Create(&payload.Cliente); res.Error != nil {
			return res.Error
		}

		payload.Empresa.ClienteID = payload.Cliente.ID
		if res := tx.Create(&payload.Empresa); res.Error != nil {
			return res.Error
		}

		payload.Solicitud.ClienteID = payload.Cliente.ID
		if res := tx.Create(&payload.Solicitud); res.Error != nil {
			return res.Error
		}

		// 🔹 Buscar operadores activos
		var operadores []auth.User
		if err := tx.Where("rol = ? AND activo = ?", "operador", true).Find(&operadores).Error; err != nil || len(operadores) == 0 {
			return fmt.Errorf("no hay operadores disponibles")
		}

		// 🔹 Calcular operador con menos asignaciones
		type Conteo struct {
			OperadorID string
			Total      int64
		}
		var conteos []Conteo
		tx.Model(&models.Asignacion{}).
			Select("operador_id, COUNT(*) as total").
			Group("operador_id").
			Scan(&conteos)

		operadorSeleccionado := operadores[0].ID
		minAsignaciones := int64(999999)
		for _, op := range operadores {
			count := int64(0)
			for _, c := range conteos {
				if c.OperadorID == op.ID {
					count = c.Total
					break
				}
			}
			if count < minAsignaciones {
				minAsignaciones = count
				operadorSeleccionado = op.ID
			}
		}

		// 🔹 Crear asignación
		asignacion := models.Asignacion{
			OperadorID:  operadorSeleccionado,
			SolicitudID: payload.Solicitud.ID,
		}
		if err := tx.Create(&asignacion).Error; err != nil {
			return fmt.Errorf("error creando asignación: %v", err)
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo registrar", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "registro exitoso",
		"cliente":   payload.Cliente,
		"empresa":   payload.Empresa,
		"solicitud": payload.Solicitud,
	})
}
