package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"url-shortener/internals/config"
	"url-shortener/internals/handlers"
	"url-shortener/internals/logger"
	"url-shortener/internals/repo"
	"url-shortener/internals/utils"
)

func main() {

	err := utils.LoadEnv(".env")
	if err != nil {
		return
	}

	postgresUrl := config.NewPostgresConfig()

	postgresRepository, _ := repo.NewPostgresRepo(postgresUrl)

	app := fiber.New()

	logger.AddLogger(app)

	cnfg := config.NewConfig()

	h := handlers.Handlers{AppConfig: cnfg, Repository: postgresRepository}

	app.Get("/:id", h.GetFullURL)

	app.Post("/shorten", h.DoShortUrlREST)

	log.Fatal(app.Listen(cnfg.ServerAddress))
}
