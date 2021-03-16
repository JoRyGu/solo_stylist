package main

import (
	"log"

	"github.com/JoRyGu/solo_stylist/configuration"
	"github.com/JoRyGu/solo_stylist/data"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	appContext, err := data.NewAppContext()
	if err != nil {
		log.Fatal("Could not acquire connection to database.")
	}

	app := fiber.New()

	configuration.ConfigureAccountsRouter(app, appContext)

	app.Listen(":3000")
}
