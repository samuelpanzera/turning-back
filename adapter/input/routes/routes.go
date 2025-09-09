package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelpanzera/turning-back/adapter/input/controller"
)

type InputPort struct {
	port int
}

func NewInputPort(port int) *InputPort {
	return &InputPort{
		port: port,
	}
}

func (ip *InputPort) InitRoutes(
	orcamentoController controller.OrcamentoControllerInterface,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/orcamentos", orcamentoController.CreateOrcamento)
		v1.GET("/orcamentos/:id", orcamentoController.FindOrcamentoByID)
		v1.GET("/orcamentos", orcamentoController.FindAllOrcamentos)
		v1.PUT("/orcamentos/:id", orcamentoController.UpdateOrcamento)
		v1.DELETE("/orcamentos/:id", orcamentoController.DeleteOrcamento)
	}

	r.POST("/orcament", orcamentoController.CreateOrcamento)

	return r
}
