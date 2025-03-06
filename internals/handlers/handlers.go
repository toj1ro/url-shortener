package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"url-shortener/internals/config"
)

var DB = map[string]string{}

type Handlers struct {
	AppConfig config.Config
}

func (h Handlers) DoShortUrl(ctx *fiber.Ctx) error {
	body := string(ctx.Body())
	shortUrlId := uuid.New().String()
	DB[shortUrlId] = body
	ctx.Status(201)
	return ctx.SendString(h.AppConfig.BaseURL + shortUrlId)
}

func (Handlers) GetFullURL(c *fiber.Ctx) error {

	id := c.Params("id")
	fullUrl := DB[id]

	c.Status(307)
	c.Set("Location", fullUrl)

	return c.SendString("")
}
