package db

import (
	"log"

	"netty/core"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

// DbSession opens a new database session.
// databaseURI should be a valid PostgreSQL connection string.
// This function will log fatal if it fails to open a connection to the
// database.
func DbSession(databaseURI string) error {
	// Create a new gorm DB instance
	var db = DB

	// Get debug mode from config
	logMode := core.Settings.DBDebugMode

	// If debug mode is enabled, set log mode to info
	if logMode {
		logLevel := logger.Info
		db, err = gorm.Open(postgres.Open(databaseURI), &gorm.Config{
			// LogMode sets the logger for gorm. Default value is silent.
			Logger: logger.Default.LogMode(logLevel),
		})
	} else {
		// Otherwise, set log mode to silent
		db, err = gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	}

	// Exit if connection fails
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	// Get underlying sql.DB and set connection pool options
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
		return err
	}

	// Set connection pool options
	sqlDB.SetMaxIdleConns(core.Settings.MaxIdleConns)
	sqlDB.SetMaxOpenConns(core.Settings.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(core.Settings.MaxLifetime)
	sqlDB.SetConnMaxIdleTime(core.Settings.MaxIdleTime)

	// Set global DB instance
	DB = db

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
