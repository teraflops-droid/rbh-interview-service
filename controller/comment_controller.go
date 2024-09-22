package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teraflops-droid/rbh-interview-service/configuration"
	"github.com/teraflops-droid/rbh-interview-service/exception"
	"github.com/teraflops-droid/rbh-interview-service/middleware"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/service"
)

type CommentController struct {
	service.CommentService
	configuration.Config
}

func NewCommentController(commentService *service.CommentService, config configuration.Config) *CommentController {
	return &CommentController{CommentService: *commentService, Config: config}
}

func (controller CommentController) Route(app *fiber.App) {
	app.Post("/v1/api/comment/create", middleware.AuthenticateJWT("USER", controller.Config), controller.CreateComment)
	app.Post("/v1/api/comment/update", middleware.AuthenticateJWT("USER", controller.Config), controller.UpdateComment)
	app.Delete("/v1/api/comment/:id", middleware.AuthenticateJWT("USER", controller.Config), controller.DeleteComment)
}

func (controller CommentController) CreateComment(c *fiber.Ctx) error {
	var request model.CommentRequest
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	username := c.Locals("username").(string)
	request.CreatedBy = username
	result, err := controller.CommentService.CreateComment(c.Context(), &request)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Success",
		Data:    result,
	})
}

func (controller CommentController) UpdateComment(c *fiber.Ctx) error {
	var request model.CommentRequest
	err := c.BodyParser(&request)
	username := c.Locals("username").(string)

	if username != request.CreatedBy {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "You are not authorized to modify this comment",
		})
	}
	exception.PanicLogging(err)
	result, err := controller.CommentService.UpdateComment(c.Context(), &request)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// DeleteComment a comment by its ID
func (controller CommentController) DeleteComment(c *fiber.Ctx) error {
	// Get the comment ID from the URL parameters
	commentId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid comment ID",
		})
	}

	// Call the service to delete the comment by ID
	err = controller.CommentService.DeleteComment(c.Context(), uint(commentId))

	// Return success message
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}
