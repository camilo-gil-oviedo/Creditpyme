package handlers

import (
	"net/http"

	"github.com/Camilo/creditPYMESbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ActualizarEstadoSolicitud(c *gin.Context) {
	var body struct {
		SolicitudID uint   `json:"solicitud_id"`
		Estado      string `json:"estado"` // "aprobado", "rechazado", "mora", "al_dia"
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Model(&models.Asignacion{}).
		Where("solicitud_id = ?", body.SolicitudID).
		Update("estado", body.Estado).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo actualizar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "estado actualizado"})
}
