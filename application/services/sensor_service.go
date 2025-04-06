package services

import (
	"Server-fuertemex/domain/entities"
	"Server-fuertemex/domain/ports"
)

type SensorService struct {
	Repo ports.SensorRepository
}

func NewSensorService(repo ports.SensorRepository) *SensorService {
	return &SensorService{Repo: repo}
}

func (s *SensorService) RegistrarEventoCaja(e entities.EventoCaja) error {
	return s.Repo.SaveEventoCaja(e)
}

func (s *SensorService) RegistrarMovimiento(m entities.Movimiento) error {
	return s.Repo.SaveMovimiento(m)
}

func (s *SensorService) RegistrarDistancia(d entities.Distancia) error {
	return s.Repo.SaveDistancia(d)
}
