package main

import (
	"fmt"
	"log"
	"os"

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

	databaseConnection := database.NewDatabaseConnection()
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

	var portInt int
	fmt.Sscanf(port, "%d", &portInt)
	return portInt
}

func initDependencies(dataBase *gorm.DB) controller.OrcamentoControllerInterface {
	orcamentoRepository := repository.NewOrcamentoRepository(dataBase)
	orcamentoService := services.NewOrcamentoUseCase(orcamentoRepository)
	return controller.NewOrcamentoControllerInterface(orcamentoService)
}
