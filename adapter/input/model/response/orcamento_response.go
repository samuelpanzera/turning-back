package response

import "time"

type OrcamentoResponse struct {
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

type CreateOrcamentoResponse struct {
	Message   string            `json:"message"`
	Orcamento OrcamentoResponse `json:"orcamento"`
}

type GetAllOrcamentosResponse struct {
	Orcamentos []OrcamentoResponse `json:"orcamentos"`
	Total      int                 `json:"total"`
}
