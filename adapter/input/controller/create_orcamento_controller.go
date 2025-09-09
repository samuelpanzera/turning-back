package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelpanzera/turning-back/adapter/input/converter"
	"github.com/samuelpanzera/turning-back/adapter/input/model/request"
	"github.com/samuelpanzera/turning-back/adapter/input/model/response"
	"github.com/samuelpanzera/turning-back/application/domain"
	"github.com/samuelpanzera/turning-back/application/port/input"
	"github.com/samuelpanzera/turning-back/configuration/logger"
	"github.com/samuelpanzera/turning-back/configuration/rest_errors"
)

func (oc *orcamentoControllerInterface) CreateOrcamento(c *gin.Context) {
	logger.Info("Init CreateOrcamento controller")

	var orcamentoRequest request.OrcamentoRequest

	if err := c.ShouldBindJSON(&orcamentoRequest); err != nil {
		logger.Error("Error trying to validate orcamento info")
		restErr := rest_errors.NewBadRequestError("Invalid request data")
		c.JSON(restErr.Code, restErr)
		return
	}

	quantidadePecas := orcamentoRequest.QuantidadePecas
	if quantidadePecas == 0 {
		quantidadePecas = orcamentoRequest.QuantidadePecasAlt
	}

	if quantidadePecas == 0 {
		logger.Error("Quantidade de peças não informada")
		restErr := rest_errors.NewBadRequestError("Quantidade de peças é obrigatória")
		c.JSON(restErr.Code, restErr)
		return
	}

	anexo := ""
	if orcamentoRequest.Anexo != nil {
		if anexoStr, ok := orcamentoRequest.Anexo.(string); ok {
			anexo = anexoStr
		}
	}

	domain := domain.OrcamentoDomain{
		Nome:            orcamentoRequest.Nome,
		Email:           orcamentoRequest.Email,
		Telefone:        orcamentoRequest.Telefone,
		QuantidadePecas: quantidadePecas,
		Descricao:       orcamentoRequest.Descricao,
		Anexo:           anexo,
	}

	domainResult, err := oc.service.CreateOrcamento(domain)
	if err != nil {
		logger.Error("Error trying to call CreateOrcamento service")
		restErr := rest_errors.NewInternalServerError("Error creating orcamento")
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("CreateOrcamento controller executed successfully")

	c.JSON(http.StatusCreated, response.CreateOrcamentoResponse{
		Message:   "Orçamento criado com sucesso",
		Orcamento: converter.ConvertDomainToResponse(domainResult),
	})
}

type orcamentoControllerInterface struct {
	service input.OrcamentoUseCaseInterface
}

func NewOrcamentoControllerInterface(service input.OrcamentoUseCaseInterface) OrcamentoControllerInterface {
	return &orcamentoControllerInterface{
		service: service,
	}
}

type OrcamentoControllerInterface interface {
	CreateOrcamento(c *gin.Context)
	FindOrcamentoByID(c *gin.Context)
	FindAllOrcamentos(c *gin.Context)
	UpdateOrcamento(c *gin.Context)
	DeleteOrcamento(c *gin.Context)
}
