package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"url-shortener/internals/config"
	"url-shortener/internals/handlers"
)

func main() {

	app := fiber.New()
	cnfg := config.NewConfig()

	h := handlers.Handlers{AppConfig: cnfg}
	app.Post("/", h.DoShortUrl)

	app.Get("/:id", h.GetFullURL)
	log.Fatal(app.Listen(cnfg.ServerAddress))
}
