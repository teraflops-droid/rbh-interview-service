package exception

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/teraflops-droid/rbh-interview-service/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Handle validation errors
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}
		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)

		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	// Handle not found errors
	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	// Handle unauthorized errors
	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
			Code:    401,
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	// General error handler
	return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}
