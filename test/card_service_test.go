package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/mocks"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/service/impl"
	"testing"
)

func TestGetCards(t *testing.T) {
	mockRepo := new(mocks.MockCardRepository)
	service := impl.NewCardServiceImpl(mockRepo)

	paginationRequest := &model.PaginationRequest{Page: 1, PageSize: 10}
	expectedCards := []entity.Card{
		{Id: 1, Title: "Card 1", Status: "Active"},
		{Id: 2, Title: "Card 2", Status: "Archived"},
		{Id: 3, Title: "Card 3", Status: "Active"},
	}

	mockRepo.On("GetAllCards", mock.Anything, paginationRequest.Page, paginationRequest.PageSize).Return(expectedCards, nil)

	cards, err := service.GetCards(context.Background(), paginationRequest)

	assert.NoError(t, err)
	assert.Len(t, *cards, 2)
	assert.Equal(t, "Card 1", (*cards)[0].Title)
	assert.Equal(t, "Card 3", (*cards)[1].Title)
	mockRepo.AssertExpectations(t)
}

func TestGetCardWithComment(t *testing.T) {
	mockRepo := new(mocks.MockCardRepository)
	service := impl.NewCardServiceImpl(mockRepo)

	expectedCard := &entity.Card{
		Id:          1,
		Title:       "Card 1",
		Description: "Description",
		Status:      "Active",
		Comments:    []entity.Comment{{Id: 1, Description: "Comment 1"}},
	}

	mockRepo.On("GetCardWithComments", mock.Anything, uint(1)).Return(expectedCard, nil)

	card, err := service.GetCardWithComment(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCard.Title, card.Title)
	assert.Equal(t, 1, len(card.Comments))
	mockRepo.AssertExpectations(t)
}

func TestCreateCard(t *testing.T) {
	mockRepo := new(mocks.MockCardRepository)
	service := impl.NewCardServiceImpl(mockRepo)

	cardRequest := &model.CardRequest{Title: "New Card", Description: "New Description", Username: "user1"}
	expectedCard := &entity.Card{
		Title:       "New Card",
		Description: "New Description",
		Status:      "Todo",
		CreatedBy:   "user1",
	}

	mockRepo.On("CreateCard", mock.Anything, mock.AnythingOfType("*entity.Card")).Return(expectedCard, nil)

	cardResponse, err := service.CreateCard(context.Background(), cardRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedCard.Title, cardResponse.Title)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCard(t *testing.T) {
	mockRepo := new(mocks.MockCardRepository)
	service := impl.NewCardServiceImpl(mockRepo)

	cardRequest := &model.CardRequest{Id: 1, Title: "Updated Card", Description: "Updated Description", Status: "Active", Username: "user1"}
	expectedCard := &entity.Card{
		Id:          1,
		Title:       "Updated Card",
		Description: "Updated Description",
		Status:      "Active",
		UpdatedBy:   "user1",
	}

	mockRepo.On("EditCard", mock.Anything, mock.AnythingOfType("*entity.Card")).Return(expectedCard, nil)

	cardResponse, err := service.UpdateCard(context.Background(), cardRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedCard.Title, cardResponse.Title)
	mockRepo.AssertExpectations(t)
}

func TestArchiveCard(t *testing.T) {
	mockRepo := new(mocks.MockCardRepository)
	service := impl.NewCardServiceImpl(mockRepo)

	mockRepo.On("UpdateCardStatus", mock.Anything, uint(1), "Archived", "user1").Return(nil)

	err := service.ArchiveCard(context.Background(), 1, "user1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
