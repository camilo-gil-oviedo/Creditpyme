package controller

import (
	"github.com/Camilo/creditPYMESbackend/auth"
	"github.com/gin-gonic/gin"
)

// AuthController maneja las rutas de autenticación
type AuthController struct {
	AuthService *auth.AuthService
}

func NewAuthController(authService *auth.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": "body inválido"})
		return
	}

	// Obtenemos el token y verificamos credenciales
	token, err := c.AuthService.Login(data.Email, data.Password)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 🔹 Obtenemos también el rol del usuario desde la base de datos
	var user auth.User
	if err := c.AuthService.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "usuario no encontrado"})
		return
	}

	// 🔹 Devolvemos token y rol
	ctx.JSON(200, gin.H{
		"token": token,
		"rol":   user.Rol,
	})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Rol      string `json:"rol"`
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": "body inválido"})
		return
	}
	if data.Rol == "" {
		data.Rol = "cliente"
	}

	token, err := c.AuthService.RegisterWithRole(data.Email, data.Password, data.Rol)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": token, "rol": data.Rol})
}

func (c *AuthController) Me(ctx *gin.Context) {
	uid, _ := ctx.Get("firebase_uid")
	claims, _ := ctx.Get("firebase_claims")
	ctx.JSON(200, gin.H{"uid": uid, "claims": claims})
}
