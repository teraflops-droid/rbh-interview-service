package service

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/model"
)

type CommentService interface {
	CreateComment(ctx context.Context, request *model.CommentRequest) (*model.CommentResponse, error)
	UpdateComment(ctx context.Context, request *model.CommentRequest) (*model.CommentResponse, error)
	DeleteComment(ctx context.Context, commentId uint) error
}
