package converter

import (
	"github.com/samuelpanzera/turning-back/adapter/output/model/entity"
	"github.com/samuelpanzera/turning-back/application/domain"
)

func ConvertDomainToEntity(orcamentoDomain domain.OrcamentoDomain) *entity.OrcamentoEntity {
	return &entity.OrcamentoEntity{
		ID:              orcamentoDomain.ID,
		CreatedAt:       orcamentoDomain.CreatedAt,
		UpdatedAt:       orcamentoDomain.UpdatedAt,
		QuantidadePecas: orcamentoDomain.QuantidadePecas,
		Descricao:       orcamentoDomain.Descricao,
		Nome:            orcamentoDomain.Nome,
		Anexo:           orcamentoDomain.Anexo,
		Email:           orcamentoDomain.Email,
		Telefone:        orcamentoDomain.Telefone,
	}
}
