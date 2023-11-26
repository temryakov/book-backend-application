package bootstrap

import (
	"fmt"
	"log"
	"review-service/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Psql(cfg *Config) *gorm.DB {

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	fmt.Printf("\nConfiguring database settings: \n Host: %s\n Port: %s\n Database Name: %s\n User: %s", cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser)

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL Database")
	}
	log.Println("Connection to database is successfully")

	client.Logger = logger.Default.LogMode(logger.Info)

	client.AutoMigrate(&domain.Review{})

	return client
}
