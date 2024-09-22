package impl

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/common/logger"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/repository"
	"github.com/teraflops-droid/rbh-interview-service/service"
	"time"
)

func NewCardServiceImpl(cardRepository *repository.CardRepository) service.CardService {
	return &cardServiceImpl{CardRepository: *cardRepository}
}

type cardServiceImpl struct {
	repository.CardRepository
}

func (cardService *cardServiceImpl) GetCards(ctx context.Context, request *model.PaginationRequest) (*[]model.CardResponse, error) {

	cards, err := cardService.CardRepository.GetAllCards(ctx, request.Page, request.PageSize)
	if err != nil {
		logger.Error(ctx, "Error while getting card")
		return nil, err
	}

	filteredCards := filterCards(cards, func(card entity.Card) bool {
		return card.Status != "Archived"
	})

	cardResponses := make([]model.CardResponse, len(filteredCards))

	// Efficiently map each card to CardResponse
	for i, card := range filteredCards {
		cardResponses[i] = mapCardToResponse(&card)
	}

	return &cardResponses, nil
}

func (cardService *cardServiceImpl) GetCardWithComment(ctx context.Context, cardId uint) (*model.CardWithCommentsResponse, error) {
	card, err := cardService.CardRepository.GetCardWithComments(ctx, cardId)
	if err != nil {
		logger.Error(ctx, "Error while getting card with comments")
		return nil, err
	}
	cardResponse := mapCardToCardWithCommentsResponse(card)

	return &cardResponse, nil
}

func (cardService *cardServiceImpl) CreateCard(ctx context.Context, card *model.CardRequest) (*model.CardResponse, error) {
	cardEntity := entity.Card{
		Title:       card.Title,
		Description: card.Description,
		Status:      "Todo",
		CreatedBy:   card.Username,
	}
	newCard, err := cardService.CardRepository.CreateCard(ctx, &cardEntity)
	if err != nil {
		logger.Error(ctx, "Error while creating card")
		return nil, err
	}
	mappedCard := mapCardToResponse(newCard)
	return &mappedCard, nil
}

func (cardService *cardServiceImpl) UpdateCard(ctx context.Context, card *model.CardRequest) (*model.CardResponse, error) {
	cardEntityToUpdate := entity.Card{
		Id:          card.Id,
		Title:       card.Title,
		Description: card.Description,
		Status:      card.Status,
		UpdatedBy:   card.Username,
	}
	updatedCard, err := cardService.CardRepository.EditCard(ctx, &cardEntityToUpdate)
	if err != nil {
		logger.Error(ctx, "Error while updating card")
		return nil, err
	}
	mappedCard := mapCardToResponse(updatedCard)
	return &mappedCard, nil
}

func mapCardToResponse(card *entity.Card) model.CardResponse {
	return model.CardResponse{
		Title:       card.Title,
		Description: card.Description,
		Status:      card.Status,
		CreatedBy:   card.CreatedBy,
		CreatedAt:   card.CreatedAt.Format(time.RFC3339), // Convert time to string
	}
}

func filterCards(cards []entity.Card, condition func(card entity.Card) bool) []entity.Card {
	var result []entity.Card
	for _, card := range cards {
		if condition(card) {
			result = append(result, card)
		}
	}
	return result
}

func mapCardToCardWithCommentsResponse(card *entity.Card) model.CardWithCommentsResponse {
	commentsResponse := make([]model.CommentResponse, len(card.Comments))
	for i, comment := range card.Comments {
		commentsResponse[i] = model.CommentResponse{
			Id:          comment.Id,
			Description: comment.Description,
			CreatedBy:   comment.CreateBy,
			CreatedAt:   comment.CreatedAt,
			UpdatedAt:   comment.UpdatedAt,
		}
	}

	return model.CardWithCommentsResponse{
		CardId:      card.Id,
		Title:       card.Title,
		Description: card.Description,
		Status:      card.Status,
		Comments:    commentsResponse,
		CreatedBy:   card.CreatedBy,
		CreatedAt:   card.CreatedAt.Format(time.RFC3339),
	}
}

func (cardService *cardServiceImpl) ArchiveCard(ctx context.Context, cardId uint, updatedBy string) error {
	err := cardService.CardRepository.UpdateCardStatus(ctx, cardId, "Archived", updatedBy)
	if err != nil {
		logger.Error(ctx, "Error while updating card")
		return err
	}

	return nil
}
