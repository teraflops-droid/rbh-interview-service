package service

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/model"
)

type CardService interface {
	CreateCard(ctx context.Context, card *model.CardRequest) (*model.CardResponse, error)
	GetCardWithComment(ctx context.Context, cardId uint) (*model.CardWithCommentsResponse, error)
	GetCards(ctx context.Context, request *model.PaginationRequest) (*[]model.CardResponse, error)
	UpdateCard(ctx context.Context, card *model.CardRequest) (*model.CardResponse, error)
	ArchiveCard(ctx context.Context, cardId uint, updatedBy string) error
}
