package repository

import (
	"port-service/internals/entity"
)

// PortRepository ...
type PortRepository struct {
	PortDB []*entity.Port
}

// GetAll is ...
func (p *PortRepository) GetAll() ([]*entity.Port, error) {
	return p.PortDB, nil
}

// Add is ...
func (p *PortRepository) Add(port *entity.Port) error {
	p.PortDB = append(p.PortDB, port)
	return nil
}

// NewPortRepository ...
func NewPortRepository() *PortRepository {
	return &PortRepository{
		PortDB: []*entity.Port{},
	}
}
