package app

func (app *App) setupRoutes() {
	app.Engine.Post("signup", app.HandleSignup())
}