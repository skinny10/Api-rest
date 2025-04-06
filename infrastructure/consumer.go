package infrastructure

import (
	"Server-fuertemex/application/services"
	"fmt"
)

// Define la estructura
type RabbitConsumer struct {
	Service *services.SensorService
}

// NewRabbitConsumer crea una nueva instancia de RabbitConsumer
func NewRabbitConsumer(s *services.SensorService) *RabbitConsumer {
	return &RabbitConsumer{Service: s}
}

// StartConsumer inicia el consumo de mensajes de RabbitMQ
func (r *RabbitConsumer) StartConsumer() {
	// aca se conecta el Rabbit para poder extraer los datoss y despues modificamos el consumer
	fmt.Println("Iniciando consumidor de RabbitMQ...")

	// Código simulado para mostrar que el consumer está funcionando
	fmt.Println("Consumidor de RabbitMQ listo para recibir mensajes")
}
