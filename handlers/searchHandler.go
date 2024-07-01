package handlers

import (
	"rest-api/ports"

	"github.com/gofiber/fiber/v2"
)

type SearchHandler struct {
	srv ports.SearchService
}

func NewSearchHandler(srv ports.SearchService) *SearchHandler {
	return &SearchHandler{
		srv: srv,
	}
}

func (searchHandler *SearchHandler) Search(c *fiber.Ctx) error {
	keyword := c.Params("keyword")
	if keyword == "" {
		return c.Status(400).JSON("Keyword is required")
	}
	keywords, err := searchHandler.srv.SearchKeywords(keyword)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(keywords)
}

func (searchHandler *SearchHandler) GetAllKeywords(c *fiber.Ctx) error {
	keywords := searchHandler.srv.GetAllKeywords()
	if keywords == nil {
		return c.Status(404).JSON("No keywords found")
	}
	return c.Status(200).JSON(keywords)
}
