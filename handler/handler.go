package handler

import (
	"github.com/brunobotter/ecommerce-usuario/config"
	"github.com/brunobotter/ecommerce-usuario/service"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMySql()
	service.InitializeService(db)
}
