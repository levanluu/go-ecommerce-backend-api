package initialize

import (
	"example.com/ecommerce/global"
	"example.com/ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
