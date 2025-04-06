package infrastructure

import (
	"Server-fuertemex/application/services"

	"github.com/gin-gonic/gin"
)

// Exporta la función para registrar rutas
func RegisterRoutes(r *gin.Engine, service *services.SensorService) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API funcionando",
		})
	})
	// Puedes añadir más rutas si es necesario
}
