package initialize

import (
	"fmt"

	"github.com/levanluu/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Host)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()

	r.Run(":8002")
}
