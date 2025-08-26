package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelpanzera/turning-back/internal/domain/entities"
	"github.com/samuelpanzera/turning-back/internal/infrastructure/config"
	"github.com/samuelpanzera/turning-back/internal/infrastructure/database"
	"github.com/samuelpanzera/turning-back/internal/infrastructure/handlers"
	"github.com/samuelpanzera/turning-back/internal/infrastructure/repositories"
	"github.com/samuelpanzera/turning-back/pkg/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config.Load()

	logger := logger.New(cfg.LogLevel, cfg.LogFormat)
	defer logger.Sync()

	db, err := database.Initialize(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to initialize database", "error", err)
	}

	if err := db.Migrate(&entities.Orcamento{}); err != nil {
		logger.Fatal("Failed to run migrations", "error", err)
	}

	if cfg.Environment == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := setupRouter(cfg, db, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Starting server", "port", port, "environment", cfg.Environment)

	if err := router.Run(":" + port); err != nil {
		logger.Fatal("Failed to start server", "error", err)
	}
}

func setupRouter(cfg *config.Config, db *database.DB, appLogger *logger.Logger) *gin.Engine {
	router := gin.New()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.Use(func(c *gin.Context) {
		appLogger.Info("Request recebido:",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"query", c.Request.URL.RawQuery,
			"user_agent", c.GetHeader("User-Agent"),
			"content_type", c.GetHeader("Content-Type"),
		)
		c.Next()
	})

	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "turning-back-api",
			"version": "1.0.0",
		})
	})

	orcamentoRepo := repositories.NewOrcamentoRepository(db.DB)
	orcamentoHandler := handlers.NewOrcamentoHandler(orcamentoRepo, appLogger)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	router.POST("/orcament", orcamentoHandler.CreateOrcamento)
	router.GET("/orcament/:id", orcamentoHandler.GetOrcamento)
	router.GET("/orcament", orcamentoHandler.GetAllOrcamentos)

	return router
}
