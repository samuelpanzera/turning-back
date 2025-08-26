package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/samuelpanzera/turning-back/internal/infrastructure/config"
	_ "modernc.org/sqlite"
)

type DB struct {
	*gorm.DB
}

func Initialize(cfg config.DatabaseConfig) (*DB, error) {
	var db *gorm.DB
	var err error

	gormLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	if cfg.Host == "sqlite" || cfg.Name == ":memory:" {
		sqlDB, err := sql.Open("sqlite", cfg.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to open sqlite database: %w", err)
		}

		db, err = gorm.Open(sqlite.Dialector{
			Conn: sqlDB,
		}, &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to sqlite database: %w", err)
		}
	} else {
		var dsn string

		if cfg.URL != "" {
			dsn = cfg.URL
		} else {
			dsn = fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
				cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode,
			)
		}

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *DB) Migrate(models ...interface{}) error {
	return db.AutoMigrate(models...)
}
