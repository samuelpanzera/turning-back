package database

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func connectSQLiteNoCGO() (*gorm.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/turning_back.db"
	}

	if err := os.MkdirAll("./data", 0755); err != nil {
		return nil, fmt.Errorf("error creating data directory: %w", err)
	}

	sqlDB, err := sql.Open("sqlite", dbPath+"?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, fmt.Errorf("error opening SQLite database: %w", err)
	}

	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("error creating GORM instance: %w", err)
	}

	return db, nil
}
