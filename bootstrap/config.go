package bootstrap

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		fmt.Print("Initializing enviroment variables")
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
	})
	return &config
}
