package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LoveSnowEx/dcard-2023-backend-intern-homework/db/page"
	"github.com/LoveSnowEx/dcard-2023-backend-intern-homework/db/pagelist"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MockConnet() (*DB, error) {
	if db != nil {
		return db, nil
	}

	// Open connection
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Error,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Migrate the schema
	conn.AutoMigrate(page.Page{}, pagelist.PageNode{}, pagelist.PageList{})

	db = &DB{DB: conn}

	return db, nil
}

func MockClose() error {
	if db == nil {
		return nil
	}

	dbConn, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	err = dbConn.Close()
	if err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	db = nil

	os.Remove("test.db")

	return nil
}
