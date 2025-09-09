package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/samuelpanzera/turning-back/adapter/output/model/entity"
	"github.com/samuelpanzera/turning-back/configuration/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var (
	database *gorm.DB
)

func NewDatabaseConnection() (*gorm.DB, error) {
	if database != nil {
		return database, nil
	}

	var err error
	dbType := strings.ToLower(os.Getenv("DB_TYPE"))

	if dbType == "" {
		env := strings.ToLower(os.Getenv("ENV"))
		switch env {
		case "production":
			dbType = "postgres"
		case "development":

			if os.Getenv("DB_HOST") != "" && os.Getenv("DB_HOST") != "sqlite" {
				dbType = "postgres"
			} else {
				dbType = "sqlite"
			}
		default:
			dbType = "sqlite"
		}
	}

	switch dbType {
	case "postgres", "postgresql":
		database, err = connectPostgreSQL()
	case "sqlite":
		database, err = connectSQLite()
	default:
		logger.Error(fmt.Sprintf("Unsupported database type: %s", dbType))
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	if err != nil {
		return nil, err
	}

	if err := database.AutoMigrate(&entity.OrcamentoEntity{}); err != nil {
		logger.Error("Error running auto migration")
		return nil, fmt.Errorf("error running auto migration: %w", err)
	}
	logger.Info("Database migration completed successfully")

	return database, nil
}

func connectPostgreSQL() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	if port == "" {
		port = "5432"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
	if timezone == "" {
		timezone = "UTC"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Error connecting to PostgreSQL database")
		return nil, err
	}

	logger.Info("Connected to PostgreSQL database successfully")
	return db, nil
}

func connectSQLite() (*gorm.DB, error) {

	db, err := connectSQLiteNoCGO()
	if err != nil {
		logger.Error("Error connecting to SQLite database")
		return nil, err
	}

	logger.Info("Connected to SQLite database successfully")
	return db, nil
}
