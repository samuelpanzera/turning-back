package repositories

import (
	"github.com/samuelpanzera/turning-back/internal/domain/entities"
	"github.com/samuelpanzera/turning-back/internal/domain/interfaces"
	"gorm.io/gorm"
)

type orcamentoRepository struct {
	db *gorm.DB
}

func NewOrcamentoRepository(db *gorm.DB) interfaces.OrcamentoRepository {
	return &orcamentoRepository{db: db}
}

func (r *orcamentoRepository) Create(orcamento *entities.Orcamento) error {
	return r.db.Create(orcamento).Error
}

func (r *orcamentoRepository) GetByID(id uint) (*entities.Orcamento, error) {
	var orcamento entities.Orcamento
	err := r.db.First(&orcamento, id).Error
	if err != nil {
		return nil, err
	}
	return &orcamento, nil
}

func (r *orcamentoRepository) GetAll() ([]*entities.Orcamento, error) {
	var orcamentos []*entities.Orcamento
	err := r.db.Find(&orcamentos).Error
	return orcamentos, err
}

func (r *orcamentoRepository) Update(orcamento *entities.Orcamento) error {
	return r.db.Save(orcamento).Error
}

func (r *orcamentoRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Orcamento{}, id).Error
}
