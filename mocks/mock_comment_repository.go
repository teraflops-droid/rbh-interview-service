package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/repository"
)

type MockCommentRepository struct {
	mock.Mock
}

// Ensure MockCommentRepository implements CommentRepository
var _ repository.CommentRepository = (*MockCommentRepository)(nil)

func (m *MockCommentRepository) CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	args := m.Called(ctx, comment)
	return args.Get(0).(entity.Comment), args.Error(1)
}

func (m *MockCommentRepository) EditComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	args := m.Called(ctx, comment)
	return args.Get(0).(*entity.Comment), args.Error(1)
}

func (m *MockCommentRepository) DeleteComment(ctx context.Context, commentId uint) error {
	args := m.Called(ctx, commentId)
	return args.Error(0)
}
