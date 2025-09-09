package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuelpanzera/turning-back/adapter/input/converter"
	"github.com/samuelpanzera/turning-back/adapter/input/model/response"
	"github.com/samuelpanzera/turning-back/configuration/logger"
	"github.com/samuelpanzera/turning-back/configuration/rest_errors"
)

func (oc *orcamentoControllerInterface) FindOrcamentoByID(c *gin.Context) {
	logger.Info("Init FindOrcamentoByID controller")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error("Error trying to validate orcamento id")
		restErr := rest_errors.NewBadRequestError("Invalid ID format")
		c.JSON(restErr.Code, restErr)
		return
	}

	domainResult, err := oc.service.FindOrcamentoByID(uint(id))
	if err != nil {
		logger.Error("Error trying to call FindOrcamentoByID service")
		restErr := rest_errors.NewNotFoundError("Orçamento não encontrado")
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("FindOrcamentoByID controller executed successfully")

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(domainResult))
}

func (oc *orcamentoControllerInterface) FindAllOrcamentos(c *gin.Context) {
	logger.Info("Init FindAllOrcamentos controller")

	domainResult, err := oc.service.FindAllOrcamentos()
	if err != nil {
		logger.Error("Error trying to call FindAllOrcamentos service")
		restErr := rest_errors.NewInternalServerError("Error fetching orcamentos")
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("FindAllOrcamentos controller executed successfully")

	orcamentosResponse := converter.ConvertDomainListToResponseList(domainResult)
	c.JSON(http.StatusOK, response.GetAllOrcamentosResponse{
		Orcamentos: orcamentosResponse,
		Total:      len(orcamentosResponse),
	})
}

func (oc *orcamentoControllerInterface) UpdateOrcamento(c *gin.Context) {
	logger.Info("Init UpdateOrcamento controller")

	c.JSON(http.StatusNotImplemented, gin.H{"message": "Update not implemented yet"})
}

func (oc *orcamentoControllerInterface) DeleteOrcamento(c *gin.Context) {
	logger.Info("Init DeleteOrcamento controller")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error("Error trying to validate orcamento id")
		restErr := rest_errors.NewBadRequestError("Invalid ID format")
		c.JSON(restErr.Code, restErr)
		return
	}

	err = oc.service.DeleteOrcamento(uint(id))
	if err != nil {
		logger.Error("Error trying to call DeleteOrcamento service")
		restErr := rest_errors.NewInternalServerError("Error deleting orcamento")
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("DeleteOrcamento controller executed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "Orçamento deletado com sucesso"})
}
