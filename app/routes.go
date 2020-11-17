package app

func (app *App) setupRoutes() {
	app.Engine.Get("login", app.HandleLogin())
}