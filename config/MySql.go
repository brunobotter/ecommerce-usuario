package config

import (
	"fmt"
	"os"

	"github.com/brunobotter/ecommerce-usuario/scheamas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeSql() (*gorm.DB, error) {
	logger := GetLogger("mysql")
	// Leia variáveis de ambiente
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Construir a string DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbName)

	//create db and connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql usuario error: %v", err)
		return nil, err
	}
	//migrate scheama
	err = db.AutoMigrate(&scheamas.Usuario{})
	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
