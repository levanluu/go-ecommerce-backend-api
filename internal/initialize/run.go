package initialize

import (
	"fmt"

	"example.com/ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitMysql()
	InitLogger()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Host, m.Username)
	global.Logger.Info("Config Log ok!!", zap.String("Ok", "Success"))
	r := InitRouter()
	r.Run(":8002")
}
