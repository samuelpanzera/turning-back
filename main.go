package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/samuelpanzera/turning-back/adapter/input/controller"
	"github.com/samuelpanzera/turning-back/adapter/input/routes"
	"github.com/samuelpanzera/turning-back/adapter/output/repository"
	"github.com/samuelpanzera/turning-back/application/services"
	"github.com/samuelpanzera/turning-back/configuration/database"
	"github.com/samuelpanzera/turning-back/configuration/logger"
	"gorm.io/gorm"
)

func main() {
	logger.Info("Starting application")

	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found")
	}

	databaseConnection, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	userController := initDependencies(databaseConnection)

	inputPort := routes.NewInputPort(getPort())
	router := inputPort.InitRoutes(userController)

	logger.Info("Server starting")
	if err := router.Run(fmt.Sprintf(":%d", getPort())); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func getPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		return 8080
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		logger.Warn(fmt.Sprintf("Invalid PORT value '%s', using default 8080", port))
		return 8080
	}
	return portInt
}

func initDependencies(dataBase *gorm.DB) controller.OrcamentoControllerInterface {
	orcamentoRepository := repository.NewOrcamentoRepository(dataBase)
	orcamentoService := services.NewOrcamentoUseCase(orcamentoRepository)
	return controller.NewOrcamentoControllerInterface(orcamentoService)
}
