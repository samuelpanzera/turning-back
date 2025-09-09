package services

import (
	"errors"

	"github.com/samuelpanzera/turning-back/application/domain"
	"github.com/samuelpanzera/turning-back/configuration/logger"
)

func (uc *orcamentoUseCase) FindOrcamentoByID(id uint) (domain.OrcamentoDomain, error) {
	logger.Info("Init FindOrcamentoByID usecase")

	return uc.orcamentoRepository.FindOrcamentoByID(id)
}

func (uc *orcamentoUseCase) FindAllOrcamentos() ([]domain.OrcamentoDomain, error) {
	logger.Info("Init FindAllOrcamentos usecase")

	return uc.orcamentoRepository.FindAllOrcamentos()
}

func (uc *orcamentoUseCase) UpdateOrcamento(orcamentoDomain domain.OrcamentoDomain) (domain.OrcamentoDomain, error) {
	logger.Info("Init UpdateOrcamento usecase")

	if !orcamentoDomain.IsValid() {
		logger.Error("Invalid orcamento data for update")
		return domain.OrcamentoDomain{}, errors.New("invalid orcamento data: nome, telefone and quantidade_pecas are required")
	}

	return uc.orcamentoRepository.UpdateOrcamento(orcamentoDomain)
}

func (uc *orcamentoUseCase) DeleteOrcamento(id uint) error {
	logger.Info("Init DeleteOrcamento usecase")

	return uc.orcamentoRepository.DeleteOrcamento(id)
}
