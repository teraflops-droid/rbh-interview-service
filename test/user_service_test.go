package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/mocks"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/service/impl"
)

func TestUserService_Authentication_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	username := "testuser"
	password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := entity.User{
		Username: username,
		Password: string(hashedPassword),
	}

	mockRepo.On("Authentication", mock.Anything, username).Return(user, nil)

	userService := impl.NewUserServiceImpl(mockRepo)

	authUser, err := userService.Authentication(context.Background(), model.UserModel{
		Username: username,
		Password: password,
	})

	assert.NoError(t, err)
	assert.Equal(t, username, authUser.Username)

	mockRepo.AssertExpectations(t)
}

func TestUserService_Authentication_InvalidPassword(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	username := "testuser"
	password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := entity.User{
		Username: username,
		Password: string(hashedPassword),
	}

	mockRepo.On("Authentication", mock.Anything, username).Return(user, nil)

	userService := impl.NewUserServiceImpl(mockRepo)

	// Test with an incorrect password
	authUser, err := userService.Authentication(context.Background(), model.UserModel{
		Username: username,
		Password: "wrongpassword",
	})

	assert.Error(t, err)
	assert.Nil(t, authUser)

	mockRepo.AssertExpectations(t)
}

func TestUserService_Register_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	username := "testuser"
	password := "password"

	userService := impl.NewUserServiceImpl(mockRepo)

	mockRepo.On("Create", mock.Anything, username, mock.AnythingOfType("string"), "USER").Return()

	err := userService.Register(context.Background(), model.UserModel{
		Username: username,
		Password: password,
	})

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
