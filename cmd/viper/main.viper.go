package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructrure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:databases`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed read config %w \n", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configration %v", err)
	}

	fmt.Println("Config Port: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("databases User: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}
