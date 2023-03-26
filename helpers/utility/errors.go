package utilityHelper

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
	"github.com/martzing/go-boilerplate/locales"
)

type CustomError struct {
	StatusCode int         `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func (e *CustomError) Error() *CustomError {
	return e
}

func (e *CustomError) SendError(c *fiber.Ctx) error {
	language := c.AcceptsLanguages("en", "th")
	statusCode := map[bool]int{true: e.StatusCode, false: fiber.StatusBadRequest}[e.StatusCode != 0]

	locales.GetLocalize(language, &e.Message)

	return c.Status(statusCode).JSON(e)
}

func ValidationError(errs validator.ValidationErrors) *CustomError {
	var data []*ErrorResponse

	for _, err := range errs {
		var e ErrorResponse

		e.FailedField = strcase.ToSnake(err.Field())
		e.Tag = err.Tag()
		e.Value = err.Param()

		data = append(data, &e)
	}
	ce := CustomError{
		Code:    "parameter-invalid",
		Message: "share.error.invalid parameter",
		Data:    data,
	}

	return &ce
}
