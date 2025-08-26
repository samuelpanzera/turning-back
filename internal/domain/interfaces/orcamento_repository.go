package interfaces

import "github.com/samuelpanzera/turning-back/internal/domain/entities"

type OrcamentoRepository interface {
	Create(orcamento *entities.Orcamento) error
	GetByID(id uint) (*entities.Orcamento, error)
	GetAll() ([]*entities.Orcamento, error)
	Update(orcamento *entities.Orcamento) error
	Delete(id uint) error
}
