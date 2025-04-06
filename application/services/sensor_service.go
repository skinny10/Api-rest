package services

import "Server-fuertemex/domain/entities"

type SensorRepository interface {
	SaveMovimiento(mov entities.Movimiento) error
	SaveDistancia(dist entities.Distancia) error
	GetAllMovimientos() ([]entities.Movimiento, error)
	GetAllDistancias() ([]entities.Distancia, error)
}

type SensorService struct {
	repo SensorRepository
}

func NewSensorService(r SensorRepository) *SensorService {
	return &SensorService{repo: r}
}

func (s *SensorService) SaveMovimiento(mov entities.Movimiento) error {
	return s.repo.SaveMovimiento(mov)
}

func (s *SensorService) SaveDistancia(dist entities.Distancia) error {
	return s.repo.SaveDistancia(dist)
}

func (s *SensorService) GetAllMovimientos() ([]entities.Movimiento, error) {
	return s.repo.GetAllMovimientos()
}

func (s *SensorService) GetAllDistancias() ([]entities.Distancia, error) {
	return s.repo.GetAllDistancias()
}
