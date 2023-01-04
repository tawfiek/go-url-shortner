package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppContext struct {
	Env Env
	DB  gorm.DB
}

type Env struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	Port       string `mapstructure:"PORT"`
}

func LoadConfig() (config Env, err error) {
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	err = viper.Unmarshal(&config)

	return
}
