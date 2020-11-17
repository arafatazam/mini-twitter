package app

import (
	"net/http"

	"github.com/arafatazam/mini-twitter/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
)

func (app *App) HandleSignup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.NewUserReq)
		if err := ctx.BodyParser(req); err != nil {
            return err
		}
		if err := validation.Validate(req); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(err)
		}
		if err := app.UserCRUD.Create(req); err != nil {
			return err
		}
		return ctx.SendStatus(http.StatusCreated)
	}
}