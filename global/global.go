package global

import (
	"example.com/ecommerce/pkg/logger"
	"example.com/ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
