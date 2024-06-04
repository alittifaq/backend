package controller

import (
	"github.com/aiteung/musik"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"ipaddress": ipaddr})
}
