package bootstrap

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBUser         string `mapstructure:"POSTGRES_USER"`
	DBPassword     string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("Can't find the file .env: %s", err)
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &config
}
