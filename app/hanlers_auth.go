package app

import "github.com/gofiber/fiber/v2"

func (app *App) HandleLogin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("testing")
	}
}