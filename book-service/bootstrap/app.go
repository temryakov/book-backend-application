package bootstrap

import (
	"github.com/IBM/sarama"
	"gorm.io/gorm"
)

type Application struct {
	Config   *Config
	DB       *gorm.DB
	Producer sarama.SyncProducer
}

func App() Application {
	app := &Application{}
	app.Config = Get()
	app.DB = Psql(app.Config)
	app.Producer = setupProducer(app.Config)
	return *app
}
