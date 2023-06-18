package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(env *Env) (*gorm.DB, error) {
	var dsn = fmt.Sprintf("host=%s user=%s password=abc123 dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", env.DBHost, env.DBUser, env.DBName, env.DBPort)
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
