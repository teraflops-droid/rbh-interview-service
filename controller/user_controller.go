package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teraflops-droid/rbh-interview-service/common"
	"github.com/teraflops-droid/rbh-interview-service/configuration"
	_ "github.com/teraflops-droid/rbh-interview-service/docs"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/service"
)

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	service.UserService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/authentication", controller.Authentication)
	app.Post("/v1/api/user/register", controller.Register)
}

// Authentication func Authenticate user.
//
//	@Description	authenticate user.
//	@Summary		authenticate user
//	@Tags			Authenticate user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserModel	true	"Request Body"
//	@Success		200		{object}	model.GeneralResponse
//	@Router			/v1/api/authentication [post]
func (controller UserController) Authentication(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request",
		})
	}
	result, err := controller.UserService.Authentication(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	role := result.UserRoles.Role
	tokenJwtResult := common.GenerateToken(result.Username, role, controller.Config)
	resultWithToken := map[string]interface{}{
		"token":    tokenJwtResult,
		"username": result.Username,
		"role":     role,
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    resultWithToken,
	})
}

// Register func Register user.
//
//	@Description	register user.
//	@Summary		register user
//	@Tags			Register user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserModel	true	"Request Body"
//	@Success		200		{object}	model.GeneralResponse
//	@Router			/v1/api/user/register [post]
func (controller UserController) Register(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request",
		})
	}
	err = controller.UserService.Register(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}
