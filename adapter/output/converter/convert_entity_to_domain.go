package converter

import (
	"github.com/samuelpanzera/turning-back/adapter/output/model/entity"
	"github.com/samuelpanzera/turning-back/application/domain"
)

func ConvertEntityToDomain(orcamentoEntity entity.OrcamentoEntity) domain.OrcamentoDomain {
	return domain.OrcamentoDomain{
		ID:              orcamentoEntity.ID,
		CreatedAt:       orcamentoEntity.CreatedAt,
		UpdatedAt:       orcamentoEntity.UpdatedAt,
		QuantidadePecas: orcamentoEntity.QuantidadePecas,
		Descricao:       orcamentoEntity.Descricao,
		Nome:            orcamentoEntity.Nome,
		Anexo:           orcamentoEntity.Anexo,
		Email:           orcamentoEntity.Email,
		Telefone:        orcamentoEntity.Telefone,
	}
}

func ConvertEntityListToDomainList(orcamentosEntity []entity.OrcamentoEntity) []domain.OrcamentoDomain {
	var orcamentosDomain []domain.OrcamentoDomain
	for _, orcamento := range orcamentosEntity {
		orcamentosDomain = append(orcamentosDomain, ConvertEntityToDomain(orcamento))
	}
	return orcamentosDomain
}
