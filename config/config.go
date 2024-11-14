package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name      string
		JwtSecret string
		Port      string
	}
	DataBase struct {
		DSN          string
		MaxIdleConns int
		MaxOpenCons  int
	}
	Redis struct {
		Addr     string
		DB       int
		Password string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file:%v", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("error into struct:%v", err)
	}
	// fmt.Printf("struct:%v", AppConfig)

	// initialzation database
	InitDB()

}
