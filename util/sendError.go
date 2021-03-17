package util

import (
	"fmt"

	"github.com/JoRyGu/solo_stylist/data/models"
	"github.com/gofiber/fiber/v2"
)

var codeMessages = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	500: "Internal Server Error",
}

func SendError(c *fiber.Ctx, statusCode int, message string) error {
	codeMessage, ok := codeMessages[statusCode]
	if ok == false {
		return fmt.Errorf("No message found for status code %d", statusCode)
	}

	errorMessage := models.HttpError{
		StatusCode:    statusCode,
		StatusMessage: codeMessage,
		Message:       message,
	}

	c.Status(statusCode)
	return c.JSON(errorMessage)
}
