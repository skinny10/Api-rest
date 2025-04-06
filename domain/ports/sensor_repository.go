package ports

import "Server-fuertemex/domain/entities"

type SensorRepository interface {
	SaveEventoCaja(event entities.EventoCaja) error
	SaveMovimiento(mov entities.Movimiento) error
	SaveDistancia(dist entities.Distancia) error
}
