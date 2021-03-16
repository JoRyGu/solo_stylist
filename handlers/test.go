package handlers

import "github.com/gofiber/fiber/v2"

type MyTestStruct struct {
	Name string
	Age  int
}

func TestHandler(c *fiber.Ctx) error {
	t := &MyTestStruct{"Josh", 34}
	return c.JSON(t)
}
