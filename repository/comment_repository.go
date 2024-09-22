package repository

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/entity"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	EditComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
	DeleteComment(ctx context.Context, commentId uint) error
}
