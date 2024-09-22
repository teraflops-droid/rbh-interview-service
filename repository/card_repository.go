package repository

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/entity"
)

type CardRepository interface {
	GetCardWithComments(ctx context.Context, cardId uint) (*entity.Card, error)
	GetAllCards(ctx context.Context, page int, pageSize int) ([]entity.Card, error)
	CreateCard(ctx context.Context, card *entity.Card) (*entity.Card, error)
	EditCard(ctx context.Context, card *entity.Card) (*entity.Card, error)
	DeleteCard(ctx context.Context, cardId string) error
	UpdateCardStatus(ctx context.Context, cardId uint, status string, updatedBy string) error
}
