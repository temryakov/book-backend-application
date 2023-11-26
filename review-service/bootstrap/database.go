package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Psql(cfg *Config) *gorm.DB {

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	fmt.Printf("\nConfiguring database settings: \n Host: %s\n Port: %s\n Database Name: %s\n User: %s", cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser)

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL Database")
	}
	return client
}
