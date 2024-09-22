package impl

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/repository"
	"gorm.io/gorm"
)

type cardRepositoryImpl struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) repository.CardRepository {
	return &cardRepositoryImpl{db: db}
}

// GetAllCards retrieve all cards
func (cardRepository *cardRepositoryImpl) GetAllCards(ctx context.Context, page int, pageSize int) ([]entity.Card, error) {
	var cards []entity.Card
	err := cardRepository.db.WithContext(ctx).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&cards).Error
	if err != nil {
		return nil, err
	}
	return cards, nil
}

// GetCardWithComments retrieves a card and its associated comments by card ID
func (cardRepository *cardRepositoryImpl) GetCardWithComments(ctx context.Context, cardId uint) (*entity.Card, error) {
	var card entity.Card

	err := cardRepository.db.WithContext(ctx).
		Preload("Comments"). // Preload the comments relationship
		Where("id = ?", cardId).
		First(&card).Error

	if err != nil {
		return nil, err
	}

	return &card, nil
}

// CreateCard creates a new card
func (cardRepository *cardRepositoryImpl) CreateCard(ctx context.Context, card *entity.Card) (*entity.Card, error) {
	err := cardRepository.db.WithContext(ctx).Create(&card).Error
	if err != nil {
		return card, err
	}
	return card, nil
}

// EditCard updates card in the database. Only fields in the card struct are updated.
func (cardRepository *cardRepositoryImpl) EditCard(ctx context.Context, card *entity.Card) (*entity.Card, error) {
	err := cardRepository.db.WithContext(ctx).Model(&entity.Card{}).Where("id = ?", card.Id).Updates(card).Error
	if err != nil {
		return card, err
	}

	err = cardRepository.db.WithContext(ctx).First(&card, card.Id).Error
	if err != nil {
		return card, err
	}

	return card, nil
}

// DeleteCard deletes
func (cardRepository *cardRepositoryImpl) DeleteCard(ctx context.Context, cardId string) error {
	err := cardRepository.db.WithContext(ctx).Delete(&entity.Card{}, cardId).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateCardStatus updates the status of a card in the database by its ID.
func (cardRepository *cardRepositoryImpl) UpdateCardStatus(ctx context.Context, cardId uint, newStatus string, updatedBy string) error {
	updates := map[string]interface{}{
		"status":     newStatus,
		"updated_by": updatedBy,
	}

	// Update the status of the card with the given ID
	err := cardRepository.db.WithContext(ctx).
		Model(&entity.Card{}).
		Where("id = ?", cardId).
		Updates(updates).Error

	if err != nil {
		return err
	}

	return nil
}
