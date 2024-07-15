package initialize

import (
	"fmt"

	"github.com/levanluu/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Host)
	InitLogger()
	global.Logger.Info("Config Log Ok", zap.String("ok", "success"))
	InitMysql()
	InitRedis()
	r := InitRouter()

	r.Run(":8002")
}
