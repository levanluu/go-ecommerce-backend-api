package global

import (
	"github.com/levanluu/go-ecommerce-backend-api/pkg/logger"
	"github.com/levanluu/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
