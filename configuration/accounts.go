package configuration

import (
	"github.com/JoRyGu/solo_stylist/data"
	"github.com/JoRyGu/solo_stylist/handlers"
	"github.com/JoRyGu/solo_stylist/services"
	"github.com/gofiber/fiber/v2"
)

func ConfigureAccountsRouter(app *fiber.App, ctx *data.AppContext) {
	a := app.Group("/accounts")
	authService := services.NewAuthService()
	accountService := services.NewAccountService(ctx, authService)

	accountsController := handlers.NewAccountController(accountService)

	a.Get("/", accountsController.GetAllAccounts)
	a.Post("/", accountsController.CreateNewAccount)
	a.Get("/:id", accountsController.GetAccountById)
}
