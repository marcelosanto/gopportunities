package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitSqlLite()
	if err != nil {
		return fmt.Errorf("Error initializing db: %v", err)
	}

	return nil
}

func GetSqLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)

	return logger
}
