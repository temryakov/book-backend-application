package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Psql(cfg *Config) *gorm.DB {

	var dsn = fmt.Sprintf("host=%s user=%s password=abc123 dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort)
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Internal database error")
	}
	return client
}
