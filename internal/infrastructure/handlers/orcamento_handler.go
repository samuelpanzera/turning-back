package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuelpanzera/turning-back/internal/domain/entities"
	"github.com/samuelpanzera/turning-back/internal/domain/interfaces"
	"github.com/samuelpanzera/turning-back/pkg/logger"
)

type OrcamentoHandler struct {
	orcamentoRepo interfaces.OrcamentoRepository
	logger        *logger.Logger
}

func NewOrcamentoHandler(orcamentoRepo interfaces.OrcamentoRepository, logger *logger.Logger) *OrcamentoHandler {
	return &OrcamentoHandler{
		orcamentoRepo: orcamentoRepo,
		logger:        logger,
	}
}


type CreateOrcamentoRequest struct {
	Nome               string      `json:"nome" binding:"required"`
	Email              string      `json:"email" binding:"omitempty,email"` 
	Telefone           string      `json:"telefone" binding:"required"`
	QuantidadePecas    int         `json:"quantidade_pecas" binding:"omitempty,min=1"`
	QuantidadePecasAlt int         `json:"quantidadePecas" binding:"omitempty,min=1"` 
	Descricao          string      `json:"descricao,omitempty"`
	Anexo              interface{} `json:"anexo,omitempty"`             
	FileUploadEnabled  bool        `json:"fileUploadEnabled,omitempty"` 
}

func (h *OrcamentoHandler) CreateOrcamento(c *gin.Context) {
	h.logger.Info("=== DEBUG: CreateOrcamento chamado ===")
	h.logger.Info("Headers recebidos:", "headers", c.Request.Header)
	h.logger.Info("Content-Type:", "content-type", c.GetHeader("Content-Type"))
	h.logger.Info("Method:", "method", c.Request.Method)

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Error("Erro ao ler body da requisição:", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao ler dados da requisição",
			"details": err.Error(),
		})
		return
	}

	h.logger.Info("JSON bruto recebido:", "raw_json", string(bodyBytes))
	h.logger.Info("Tamanho do body:", "size", len(bodyBytes))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var request CreateOrcamentoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Error("Erro no bind JSON:", "error", err.Error())
		h.logger.Error("JSON que causou erro:", "json", string(bodyBytes))

		c.JSON(http.StatusBadRequest, gin.H{
			"error":         "Invalid request data",
			"details":       err.Error(),
			"received_json": string(bodyBytes),
			"help":          "Campos obrigatórios: nome, telefone, quantidade_pecas (ou quantidadePecas). Email é opcional.",
		})
		return
	}

	
	requestJSON, _ := json.Marshal(request)
	h.logger.Info("Request parseado com sucesso:", "request", string(requestJSON))

	
	orcamento := entities.Orcamento{
		Nome:      request.Nome,
		Email:     request.Email,
		Telefone:  request.Telefone,
		Descricao: request.Descricao,
	}

	
	if request.QuantidadePecas > 0 {
		orcamento.QuantidadePecas = request.QuantidadePecas
	} else if request.QuantidadePecasAlt > 0 {
		orcamento.QuantidadePecas = request.QuantidadePecasAlt
	} else {
		h.logger.Error("Quantidade de peças não informada")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Quantidade de peças é obrigatória",
			"help":  "Use 'quantidade_pecas' ou 'quantidadePecas' com valor maior que 0",
		})
		return
	}

	
	if request.Anexo != nil {
		if anexoStr, ok := request.Anexo.(string); ok {
			orcamento.Anexo = anexoStr
		}
	}

	
	if orcamento.Email == "" {
		h.logger.Info("Email não fornecido - campo opcional")
	} else {
		h.logger.Info("Email fornecido:", "email", orcamento.Email)
	}

	
	orcamentoJSON, _ := json.Marshal(orcamento)
	h.logger.Info("Entidade criada:", "orcamento", string(orcamentoJSON))

	if err := h.orcamentoRepo.Create(&orcamento); err != nil {
		h.logger.Error("Erro ao criar orçamento no banco:", "error", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create orcamento",
			"details": err.Error(),
		})
		return
	}

	h.logger.Info("Orçamento criado com sucesso:", "id", orcamento.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Orçamento criado com sucesso",
		"orcamento": orcamento,
	})
}

func (h *OrcamentoHandler) GetOrcamento(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	orcamento, err := h.orcamentoRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Orçamento não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, orcamento)
}

func (h *OrcamentoHandler) GetAllOrcamentos(c *gin.Context) {
	orcamentos, err := h.orcamentoRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch orcamentos",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orcamentos": orcamentos,
		"total":      len(orcamentos),
	})
}
