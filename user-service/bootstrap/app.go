package bootstrap

import "gorm.io/gorm"

type Application struct {
	Config *Config
	DB     *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Config = Get()
	app.DB = Psql(app.Config)
	return *app
}
