package domain

import (
	"time"
)

type OrcamentoDomain struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	QuantidadePecas int       `json:"quantidade_pecas"`
	Descricao       string    `json:"descricao,omitempty"`
	Nome            string    `json:"nome"`
	Anexo           string    `json:"anexo,omitempty"`
	Email           string    `json:"email,omitempty"`
	Telefone        string    `json:"telefone"`
}

func (o *OrcamentoDomain) IsValid() bool {
	return o.Nome != "" && o.Telefone != "" && o.QuantidadePecas > 0
}
