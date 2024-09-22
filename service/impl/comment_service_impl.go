package impl

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/common/logger"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/repository"
	"github.com/teraflops-droid/rbh-interview-service/service"
)

type commentServiceImpl struct {
	repository.CommentRepository
}

func NewCommentServiceImpl(commentRepository repository.CommentRepository) service.CommentService {
	return commentServiceImpl{CommentRepository: commentRepository}
}

func (c commentServiceImpl) UpdateComment(ctx context.Context, request *model.CommentRequest) (*model.CommentResponse, error) {
	commentEntity := entity.Comment{
		Id:          request.Id,
		Description: request.Description,
		CardId:      request.CardId,
		CreateBy:    request.CreatedBy,
	}
	result, err := c.CommentRepository.EditComment(ctx, &commentEntity)
	if err != nil {
		logger.Error(ctx, "Error while editing card")
		return nil, err
	}
	mappedResult := mapCommentToResponse(result)
	return &mappedResult, nil
}

func (c commentServiceImpl) DeleteComment(ctx context.Context, commentId uint) error {
	err := c.CommentRepository.DeleteComment(ctx, commentId)
	return err
}

func (c commentServiceImpl) CreateComment(ctx context.Context, request *model.CommentRequest) (*model.CommentResponse, error) {
	commentEntity := entity.Comment{
		Description: request.Description,
		CardId:      request.CardId,
		CreateBy:    request.CreatedBy,
	}
	result, err := c.CommentRepository.CreateComment(ctx, commentEntity)
	if err != nil {
		logger.Error(ctx, "Error while creating comment")
		return nil, err
	}

	mappedResult := mapCommentToResponse(&result)
	return &mappedResult, nil
}

func mapCommentToResponse(card *entity.Comment) model.CommentResponse {
	return model.CommentResponse{
		Id:          card.Id,
		Description: card.Description,
		CreatedBy:   card.CreateBy,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
	}
}
