package impl

import (
	"context"
	"errors"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/repository"
	"github.com/teraflops-droid/rbh-interview-service/service"

	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (userService *userServiceImpl) Authentication(ctx context.Context, model model.UserModel) (*entity.User, error) {
	userResult, err := userService.UserRepository.Authentication(ctx, model.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(model.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &userResult, nil
}

func (userService *userServiceImpl) Register(ctx context.Context, model model.UserModel) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userService.UserRepository.Create(ctx, model.Username, string(hashPassword), "USER")
	return nil
}
