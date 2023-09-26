package bootstrap

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBUser         string `mapstructure:"POSTGRES_USER"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	SecretKey      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	ExpiryHours    int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOURS"`
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
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &config
}
