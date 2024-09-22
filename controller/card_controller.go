package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teraflops-droid/rbh-interview-service/configuration"
	"github.com/teraflops-droid/rbh-interview-service/middleware"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/service"
)

type CardController struct {
	service.CardService
	configuration.Config
}

func NewCardController(cardService *service.CardService, config configuration.Config) *CardController {
	return &CardController{CardService: *cardService, Config: config}
}

func (controller CardController) Route(app *fiber.App) {
	app.Post("/v1/api/card/create", middleware.AuthenticateJWT("USER", controller.Config), controller.CreateCard)
	app.Get("/v1/api/cards", middleware.AuthenticateJWT("USER", controller.Config), controller.GetCards)
	app.Get("/v1/api/card/:id", middleware.AuthenticateJWT("USER", controller.Config), controller.GetCardWithComments)
	app.Put("/v1/api/card/update", middleware.AuthenticateJWT("USER", controller.Config), controller.UpdateCard)
	app.Patch("/v1/api/card/:id/archive", middleware.AuthenticateJWT("USER", controller.Config), controller.ArchiveCard)
}

// CreateCard func Create card.
//
//	@Description	create card.
//	@Summary		create new card
//	@Tags			Create user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.CardRequest	true	"Request Body"
//	@Success		200		{object}	model.GeneralResponse
//	@Router			/v1/api/card/create [post]
func (controller CardController) CreateCard(c *fiber.Ctx) error {
	var request model.CardRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request",
		})
	}
	username := c.Locals("username").(string)
	request.Username = username
	result, err := controller.CardService.CreateCard(c.Context(), &request)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Success",
		Data:    result,
	})
}

// GetCardWithComments func Get card with comments.
//
//	@Description	get card with comments.
//	@Summary		get card with comments
//	@Tags			Get card
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Card ID"
//	@Success		200	{object}	model.GeneralResponse
//	@Router			/v1/api/card/:id [get]
func (controller CardController) GetCardWithComments(c *fiber.Ctx) error {
	cardId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request",
		})
	}
	result, err := controller.CardService.GetCardWithComment(c.Context(), uint(cardId))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// GetCards retrieves a paginated list of cards.
//
//	@Description	Get a list of cards with pagination.
//	@Summary		Get paginated cards
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"Page number"				default(1)
//	@Param			pageSize	query		int	false	"Number of items per page"	default(10)
//	@Success		200			{object}	model.GeneralResponse
//	@Failure		400			{object}	model.GeneralResponse	"Bad Request"
//	@Router			/v1/api/cards [get]
func (controller CardController) GetCards(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	request := model.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
	}
	result, err := controller.CardService.GetCards(c.Context(), &request)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// UpdateCard updates a card.
//
//	@Description	Update a card's details.
//	@Summary		Update card
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.CardRequest	true	"Card update request body"
//	@Success		200		{object}	model.GeneralResponse
//	@Failure		400		{object}	model.GeneralResponse	"Bad Request"
//	@Failure		500		{object}	model.GeneralResponse	"Internal Server Error"
//	@Router			/v1/api/card/update [put]
func (controller CardController) UpdateCard(c *fiber.Ctx) error {
	var request model.CardRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request",
		})
	}
	username := c.Locals("username").(string)
	request.Username = username

	updatedCard, err := controller.CardService.UpdateCard(c.Context(), &request)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    updatedCard,
	})
}

// ArchiveCard archives a card.
//
//	@Description	Archive a card by its ID.
//	@Summary		Archive card
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Card ID"
//	@Success		200	{object}	model.GeneralResponse
//	@Failure		400	{object}	model.GeneralResponse	"Bad Request"
//	@Failure		404	{object}	model.GeneralResponse	"Card not found"
//	@Router			/v1/api/cards/{id}/archive [patch]
func (controller CardController) ArchiveCard(c *fiber.Ctx) error {
	cardId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request",
		})
	}
	username := c.Locals("username").(string)

	err = controller.CardService.ArchiveCard(c.Context(), uint(cardId), username)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})

}
