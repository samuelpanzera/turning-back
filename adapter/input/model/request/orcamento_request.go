package request

type OrcamentoRequest struct {
	Nome               string      `json:"nome" binding:"required"`
	Email              string      `json:"email" binding:"omitempty,email"`
	Telefone           string      `json:"telefone" binding:"required"`
	QuantidadePecas    int         `json:"quantidade_pecas" binding:"omitempty,min=1"`
	QuantidadePecasAlt int         `json:"quantidadePecas" binding:"omitempty,min=1"`
	Descricao          string      `json:"descricao,omitempty"`
	Anexo              interface{} `json:"anexo,omitempty"`
	FileUploadEnabled  bool        `json:"fileUploadEnabled,omitempty"`
}

type UpdateOrcamentoRequest struct {
	ID              uint   `json:"id" binding:"required"`
	Nome            string `json:"nome" binding:"required"`
	Email           string `json:"email" binding:"omitempty,email"`
	Telefone        string `json:"telefone" binding:"required"`
	QuantidadePecas int    `json:"quantidade_pecas" binding:"required,min=1"`
	Descricao       string `json:"descricao,omitempty"`
	Anexo           string `json:"anexo,omitempty"`
}
