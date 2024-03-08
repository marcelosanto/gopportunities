package config

import (
	"com.github/marcelosanto/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlLite(path string) (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
	
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite AutoMigrate error: %v", err)
		return nil, err
	}

	return db, nil
}