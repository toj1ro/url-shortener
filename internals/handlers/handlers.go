package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"url-shortener/internals/config"
	"url-shortener/internals/repo"
)

type Handlers struct {
	AppConfig  config.Config
	Repository repo.PostgresRepo
}

type fullUrl struct {
	Url string `json:"url"`
}

func (h Handlers) DoShortUrlREST(ctx *fiber.Ctx) error {
	body := ctx.Body()
	var full fullUrl

	err := json.Unmarshal(body, &full)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Json invalid!")
	}

	shortUrl, _ := h.Repository.Create(full.Url)

	ctx.Status(201)
	return ctx.SendString(h.AppConfig.BaseURL + shortUrl)
}

func (h Handlers) GetFullURL(c *fiber.Ctx) error {

	id := c.Params("id")
	fullUrl, _ := h.Repository.Get(id)

	c.Status(307)
	c.Set("Location", fullUrl)

	return c.SendString("")
}
