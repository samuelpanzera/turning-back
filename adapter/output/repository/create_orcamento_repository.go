package repository

import (
	"github.com/samuelpanzera/turning-back/adapter/output/converter"
	"github.com/samuelpanzera/turning-back/application/domain"
	"github.com/samuelpanzera/turning-back/configuration/logger"
)

func (or *orcamentoRepository) CreateOrcamento(orcamentoDomain domain.OrcamentoDomain) (domain.OrcamentoDomain, error) {
	logger.Info("Init CreateOrcamento repository")

	orcamentoEntity := converter.ConvertDomainToEntity(orcamentoDomain)

	if err := or.databaseConnection.Create(orcamentoEntity).Error; err != nil {
		logger.Error("Error trying to create orcamento")
		return domain.OrcamentoDomain{}, err
	}

	logger.Info("CreateOrcamento repository executed successfully")
	return converter.ConvertEntityToDomain(*orcamentoEntity), nil
}
