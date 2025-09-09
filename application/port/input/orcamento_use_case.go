package input

import "github.com/samuelpanzera/turning-back/application/domain"

type OrcamentoUseCaseInterface interface {
	CreateOrcamento(orcamentoDomain domain.OrcamentoDomain) (domain.OrcamentoDomain, error)
	FindOrcamentoByID(id uint) (domain.OrcamentoDomain, error)
	FindAllOrcamentos() ([]domain.OrcamentoDomain, error)
	UpdateOrcamento(orcamentoDomain domain.OrcamentoDomain) (domain.OrcamentoDomain, error)
	DeleteOrcamento(id uint) error
}
