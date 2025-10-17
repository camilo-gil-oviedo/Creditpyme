package server

import "github.com/gin-gonic/gin"

// Start arranca el servidor Gin en la dirección indicada
func Start(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
