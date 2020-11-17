package app

import (
	"fmt"

	"github.com/arafatazam/mini-twitter/service"
	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct{
	Config *Config
	DB *gorm.DB
	UserCRUD *service.UserCRUD
	Engine *fiber.App
}

func (app *App) readConfig() {
	app.Config = new(Config)
	env.Parse(app.Config)
}

func (app *App) setupDB() {
	c := app.Config
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
				c.DBUser, c.DBPass, c.DBName, c.DBPort)
	app.DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func (app *App) injectServices(){
	app.UserCRUD = &service.UserCRUD{
		DB: app.DB,
	}
}

func (app *App) setupEngine() {
	app.Engine = fiber.New()
	app.Engine.Use(recover.New())
	app.setupRoutes()
}

func (app *App) Run() error {
	return app.Engine.Listen(":"+app.Config.Port)
}

func New() *App {
	app := new(App)
	app.readConfig()
	app.setupDB()
	app.injectServices()
	app.setupEngine()
	return app
}