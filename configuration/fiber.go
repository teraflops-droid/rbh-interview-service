package configuration

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teraflops-droid/rbh-interview-service/exception"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
