package service

import (
	"port-service/internals/entity"
	"port-service/internals/repository"
)

// PortService is ...
type PortService struct {
	PortRepository *repository.PortRepository
}

// GetPort ...
func (p *PortService) GetPort() ([]*entity.Port, error) {
	port, err := p.PortRepository.GetAll()
	if err != nil {
		return []*entity.Port{}, err
	}
	return port, nil
}

// AddPort ...
func (p *PortService) AddPort(port *entity.Port) error {
	err := p.PortRepository.Add(port)
	if err != nil {
		return err
	}
	return nil
}

// NewPortService is ...
func NewPortService() *PortService {
	return &PortService{
		PortRepository: repository.NewPortRepository(),
	}
}
