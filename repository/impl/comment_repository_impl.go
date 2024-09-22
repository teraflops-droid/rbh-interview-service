package impl

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/repository"
	"gorm.io/gorm"
)

type commentRepositoryImpl struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepositoryImpl{db: db}
}

// CreateComment creates a new card
func (commentRepository *commentRepositoryImpl) CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	err := commentRepository.db.WithContext(ctx).Create(&comment).Error
	if err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}

// EditComment updates comment in database.
func (commentRepository *commentRepositoryImpl) EditComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	err := commentRepository.db.WithContext(ctx).Model(&entity.Comment{}).Where("id = ?", comment.Id).Updates(comment).Error
	if err != nil {
		return nil, err
	}

	err = commentRepository.db.WithContext(ctx).First(&comment, comment.Id).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// DeleteComment deletes a comment from the database by its ID.
func (commentRepository *commentRepositoryImpl) DeleteComment(ctx context.Context, commentId uint) error {
	err := commentRepository.db.WithContext(ctx).Delete(&entity.Comment{}, commentId).Error
	if err != nil {
		return err
	}
	return nil
}
