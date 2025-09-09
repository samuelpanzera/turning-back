package entity

import (
	"time"

	"gorm.io/gorm"
)

type OrcamentoEntity struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	QuantidadePecas int    `json:"quantidade_pecas" gorm:"not null"`
	Descricao       string `json:"descricao,omitempty"`
	Nome            string `json:"nome" gorm:"not null"`
	Anexo           string `json:"anexo,omitempty"`
	Email           string `json:"email,omitempty"`
	Telefone        string `json:"telefone" gorm:"not null"`
}

func (OrcamentoEntity) TableName() string {
	return "orcamentos"
}
