package mocks

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teraflops-droid/rbh-interview-service/entity"
)

type MockUserRepository struct {
	mock.Mock
}

// Create simulates creating a new user
func (m *MockUserRepository) Create(ctx context.Context, username string, password string, role string) {
	m.Called(ctx, username, password, role)
}

// Authentication simulates user authentication
func (m *MockUserRepository) Authentication(ctx context.Context, username string) (entity.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(entity.User), args.Error(1)
}

// Example of a helper function to create a user for testing
func NewMockUser(username string) entity.User {
	return entity.User{
		Username: username,
		Password: "hashed_password", // Replace with the actual hashed password if needed
		IsActive: true,
		UserRoles: entity.UserRole{
			Id:       uuid.New(),
			Username: username,
			Role:     "user_role", // Replace with actual roles for testing
		},
	}
}
