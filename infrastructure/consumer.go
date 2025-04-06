package infrastructure

import (
	"Server-fuertemex/application/services"
	"fmt"
)

type RabbitConsumer struct {
	Service *services.SensorService
}

func NewRabbitConsumer(s *services.SensorService) *RabbitConsumer {
	return &RabbitConsumer{Service: s}
}

func (r *RabbitConsumer) StartConsumer() {
	fmt.Println("Iniciando consumidor de RabbitMQ...")
	fmt.Println("Consumidor de RabbitMQ listo para recibir mensajes")
}
