package converter

import (
	"github.com/samuelpanzera/turning-back/adapter/input/model/response"
	"github.com/samuelpanzera/turning-back/application/domain"
)

func ConvertDomainToResponse(orcamentoDomain domain.OrcamentoDomain) response.OrcamentoResponse {
	return response.OrcamentoResponse{
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

func ConvertDomainListToResponseList(orcamentosDomain []domain.OrcamentoDomain) []response.OrcamentoResponse {
	var orcamentosResponse []response.OrcamentoResponse
	for _, orcamento := range orcamentosDomain {
		orcamentosResponse = append(orcamentosResponse, ConvertDomainToResponse(orcamento))
	}
	return orcamentosResponse
}
