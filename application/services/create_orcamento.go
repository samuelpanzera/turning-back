package services

import (
	"errors"

	"github.com/samuelpanzera/turning-back/application/domain"
	"github.com/samuelpanzera/turning-back/application/port/output"
	"github.com/samuelpanzera/turning-back/configuration/logger"
)

func (uc *orcamentoUseCase) CreateOrcamento(orcamentoDomain domain.OrcamentoDomain) (domain.OrcamentoDomain, error) {
	logger.Info("Init CreateOrcamento usecase")

	if !orcamentoDomain.IsValid() {
		logger.Error("Invalid orcamento data")
		return domain.OrcamentoDomain{}, errors.New("invalid orcamento data: nome, telefone and quantidade_pecas are required")
	}

	orcamentoDomainRepository, err := uc.orcamentoRepository.CreateOrcamento(orcamentoDomain)
	if err != nil {
		logger.Error("Error trying to call repository")
		return domain.OrcamentoDomain{}, err
	}

	logger.Info("CreateOrcamento usecase executed successfully")
	return orcamentoDomainRepository, nil
}

type orcamentoUseCase struct {
	orcamentoRepository output.OrcamentoRepositoryInterface
}

func NewOrcamentoUseCase(orcamentoRepository output.OrcamentoRepositoryInterface) *orcamentoUseCase {
	return &orcamentoUseCase{
		orcamentoRepository: orcamentoRepository,
	}
}
