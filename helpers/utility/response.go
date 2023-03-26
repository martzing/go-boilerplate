package utilityHelper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/martzing/go-boilerplate/locales"
)

type Response struct {
	StatusCode int         `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func (r *Response) Send(c *fiber.Ctx) error {
	language := c.AcceptsLanguages("en", "th")
	statusCode := map[bool]int{true: r.StatusCode, false: fiber.StatusOK}[r.StatusCode != 0]

	locales.GetLocalize(language, &r.Message)

	return c.Status(statusCode).JSON(r)
}
