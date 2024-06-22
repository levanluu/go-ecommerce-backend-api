package initialize

import (
	"fmt"

	"github.com/levanluu/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed read configution %w \n", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configration %v", err)
	}
}
