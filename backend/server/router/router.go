package router

import (
	"net/http"

	firebaseauth "firebase.google.com/go/v4/auth"
	"github.com/Camilo/creditPYMESbackend/server/controller"
	"github.com/Camilo/creditPYMESbackend/server/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(authController *controller.AuthController, authClient *firebaseauth.Client, dbConn *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)

	// Health endpoint
	r.GET("/health", func(c *gin.Context) {
		// Check DB
		sqlDB, err := dbConn.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"db": "error", "error": err.Error()})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"db": "unreachable", "error": err.Error()})
			return
		}

		// Check Firebase client
		if authClient == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"firebase": "not initialized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthMiddleware(authClient))
	{
		authGroup.GET("/me", authController.Me)
	}

	return r
}
