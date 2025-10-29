package handlers

import (
	"net/http"

	"github.com/Camilo/creditPYMESbackend/auth"
	"github.com/Camilo/creditPYMESbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EstadisticasAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var totalOperadores int64
	var totalClientes int64
	var creditosMora int64
	var creditosAlDia int64

	db.Model(&models.Asignacion{}).Where("estado = ?", "mora").Count(&creditosMora)
	db.Model(&models.Asignacion{}).Where("estado = ?", "al_dia").Count(&creditosAlDia)
	db.Model(&models.Cliente{}).Count(&totalClientes)
	db.Model(&auth.User{}).Where("rol = ?", "operador").Count(&totalOperadores)

	c.JSON(http.StatusOK, gin.H{
		"total_operadores": totalOperadores,
		"total_clientes":   totalClientes,
		"creditos_mora":    creditosMora,
		"creditos_al_dia":  creditosAlDia,
	})
}
