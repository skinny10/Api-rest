package main

import (
	"Server-fuertemex/application/services"
	"Server-fuertemex/domain/entities"
	"Server-fuertemex/infrastructure"
	"fmt"

	"github.com/gin-gonic/gin"
)

// MockRepository es una implementación de memoria para SensorRepository
type MockRepository struct{}

// Implementación de los métodos de la interfaz ports.SensorRepository
func (m *MockRepository) SaveEventoCaja(event entities.EventoCaja) error {
	fmt.Printf("Mock: Guardando evento de caja %v\n", event)
	return nil
}

func (m *MockRepository) SaveMovimiento(mov entities.Movimiento) error {
	fmt.Printf("Mock: Guardando movimiento %v\n", mov)
	return nil
}

func (m *MockRepository) SaveDistancia(dist entities.Distancia) error {
	fmt.Printf("Mock: Guardando distancia %v\n", dist)
	return nil
}

func main() {
	// Crear un repositorio mock en lugar de MySQL
	repo := &MockRepository{}

	// Crear el servicio usando el repositorio mock
	service := services.NewSensorService(repo)

	// Iniciar RabbitMQ en una goroutine
	go infrastructure.NewRabbitConsumer(service).StartConsumer()

	// Configurar la API REST
	r := gin.Default()
	infrastructure.RegisterRoutes(r, service)
	r.Run(":8080")
}
