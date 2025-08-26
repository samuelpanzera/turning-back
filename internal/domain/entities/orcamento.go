package entities

import (
	"time"

	"gorm.io/gorm"
)

type Orcamento struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	QuantidadePecas int    `json:"quantidade_pecas" gorm:"not null" binding:"required,min=1"`
	Descricao       string `json:"descricao,omitempty"`
	Nome            string `json:"nome" gorm:"not null" binding:"required"`
	Anexo           string `json:"anexo,omitempty"`
	Email           string `json:"email,omitempty" gorm:"" binding:"omitempty,email"`
	Telefone        string `json:"telefone" gorm:"not null" binding:"required"`
}

func (Orcamento) TableName() string {
	return "orcamentos"
}
