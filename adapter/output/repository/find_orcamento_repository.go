package repository

import (
	"errors"

	"github.com/samuelpanzera/turning-back/adapter/output/converter"
	"github.com/samuelpanzera/turning-back/adapter/output/model/entity"
	"github.com/samuelpanzera/turning-back/application/domain"
	"github.com/samuelpanzera/turning-back/application/port/output"
	"github.com/samuelpanzera/turning-back/configuration/logger"
	"gorm.io/gorm"
)

func (or *orcamentoRepository) FindOrcamentoByID(id uint) (domain.OrcamentoDomain, error) {
	logger.Info("Init FindOrcamentoByID repository")

	var orcamentoEntity entity.OrcamentoEntity
	if err := or.databaseConnection.First(&orcamentoEntity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("Orcamento not found")
			return domain.OrcamentoDomain{}, errors.New("orcamento not found")
		}
		logger.Error("Error trying to find orcamento")
		return domain.OrcamentoDomain{}, err
	}

	logger.Info("FindOrcamentoByID repository executed successfully")
	return converter.ConvertEntityToDomain(orcamentoEntity), nil
}

func (or *orcamentoRepository) FindAllOrcamentos() ([]domain.OrcamentoDomain, error) {
	logger.Info("Init FindAllOrcamentos repository")

	var orcamentosEntity []entity.OrcamentoEntity
	if err := or.databaseConnection.Find(&orcamentosEntity).Error; err != nil {
		logger.Error("Error trying to find all orcamentos")
		return nil, err
	}

	logger.Info("FindAllOrcamentos repository executed successfully")
	return converter.ConvertEntityListToDomainList(orcamentosEntity), nil
}

func (or *orcamentoRepository) UpdateOrcamento(orcamentoDomain domain.OrcamentoDomain) (domain.OrcamentoDomain, error) {
	logger.Info("Init UpdateOrcamento repository")

	orcamentoEntity := converter.ConvertDomainToEntity(orcamentoDomain)

	if err := or.databaseConnection.Save(orcamentoEntity).Error; err != nil {
		logger.Error("Error trying to update orcamento")
		return domain.OrcamentoDomain{}, err
	}

	logger.Info("UpdateOrcamento repository executed successfully")
	return converter.ConvertEntityToDomain(*orcamentoEntity), nil
}

func (or *orcamentoRepository) DeleteOrcamento(id uint) error {
	logger.Info("Init DeleteOrcamento repository")

	if err := or.databaseConnection.Delete(&entity.OrcamentoEntity{}, id).Error; err != nil {
		logger.Error("Error trying to delete orcamento")
		return err
	}

	logger.Info("DeleteOrcamento repository executed successfully")
	return nil
}

type orcamentoRepository struct {
	databaseConnection *gorm.DB
}

func NewOrcamentoRepository(databaseConnection *gorm.DB) output.OrcamentoRepositoryInterface {
	return &orcamentoRepository{
		databaseConnection: databaseConnection,
	}
}
