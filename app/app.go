package app

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct{
	Engine *fiber.App
	DB *gorm.DB
}

func (app *App) readConfig() {
	// Todo read in config here
}

func (app *App) setupDB() {
	// Todo
}

func (app *App) setupEngine() {
	app.Engine = fiber.New()
	app.setupRoutes()
}

func (app *App) Run() error {
	return app.Engine.Listen(":8080")
}

func New() *App {
	app := new(App)
	app.readConfig()
	app.setupDB()
	app.setupEngine()
	return app
}