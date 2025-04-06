package main

import (
	"Server-fuertemex/application/services"
	"Server-fuertemex/domain/entities"
	"Server-fuertemex/infrastructure"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ----- Mock Repositorio -----
type MockRepository struct{}

func (m *MockRepository) SaveMovimiento(mov entities.Movimiento) error {
	return nil
}

func (m *MockRepository) SaveDistancia(dist entities.Distancia) error {
	fmt.Printf("Mock: Guardando distancia %v\n", dist)
	return nil
}

func (m *MockRepository) GetAllMovimientos() ([]entities.Movimiento, error) {
	return []entities.Movimiento{
		{ID: 1, Descripcion: "Movimiento 1"},
	}, nil
}

func (m *MockRepository) GetAllDistancias() ([]entities.Distancia, error) {
	return []entities.Distancia{
		{ID: 1, Valor: 123.45},
	}, nil
}

func main() {
	repo := &MockRepository{}
	service := services.NewSensorService(repo)

	// Iniciar RabbitMQ en una goroutine
	go infrastructure.NewRabbitConsumer(service).StartConsumer()

	// Configurar la API REST
	r := gin.Default()

	r.GET("/movimientos", func(c *gin.Context) {
		movimientos, err := service.GetAllMovimientos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, movimientos)
	})

	r.GET("/distancias", func(c *gin.Context) {
		distancias, err := service.GetAllDistancias()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, distancias)
	})

	r.Run(":8080")
}
