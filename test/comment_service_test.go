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

func TestCreateComment(t *testing.T) {
	mockRepo := new(mocks.MockCommentRepository)
	commentService := impl.NewCommentServiceImpl(mockRepo)

	request := &model.CommentRequest{
		Description: "Test Comment",
		CardId:      1,
		CreatedBy:   "user1",
	}

	expectedComment := entity.Comment{
		Id:          1,
		Description: request.Description,
		CardId:      request.CardId,
		CreateBy:    request.CreatedBy,
	}

	mockRepo.On("CreateComment", mock.Anything, mock.AnythingOfType("entity.Comment")).Return(expectedComment, nil)

	result, err := commentService.CreateComment(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, expectedComment.Id, result.Id)
	assert.Equal(t, expectedComment.Description, result.Description)
	assert.Equal(t, expectedComment.CreateBy, result.CreatedBy)

	mockRepo.AssertExpectations(t)
}

func TestUpdateComment(t *testing.T) {
	mockRepo := new(mocks.MockCommentRepository)
	commentService := impl.NewCommentServiceImpl(mockRepo)

	request := &model.CommentRequest{
		Id:          1,
		Description: "Updated Comment",
		CardId:      1,
		CreatedBy:   "user1",
	}

	expectedComment := entity.Comment{
		Id:          request.Id,
		Description: request.Description,
		CardId:      request.CardId,
		CreateBy:    request.CreatedBy,
	}

	mockRepo.On("EditComment", mock.Anything, &expectedComment).Return(&expectedComment, nil)

	result, err := commentService.UpdateComment(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, expectedComment.Id, result.Id)
	assert.Equal(t, expectedComment.Description, result.Description)
	assert.Equal(t, expectedComment.CreateBy, result.CreatedBy)

	mockRepo.AssertExpectations(t)
}

func TestDeleteComment(t *testing.T) {
	mockRepo := new(mocks.MockCommentRepository)
	commentService := impl.NewCommentServiceImpl(mockRepo)

	commentId := uint(1)

	mockRepo.On("DeleteComment", mock.Anything, commentId).Return(nil)

	err := commentService.DeleteComment(context.Background(), commentId)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
