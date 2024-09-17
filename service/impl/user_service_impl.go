package impl

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/exception"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/repository"
	"github.com/teraflops-droid/rbh-interview-service/service"

	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (userService *userServiceImpl) Authentication(ctx context.Context, model model.UserModel) entity.User {
	userResult, err := userService.UserRepository.Authentication(ctx, model.Username)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(model.Password))
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "incorrect username and password",
		})
	}
	return userResult
}

func (userService *userServiceImpl) Register(model model.UserModel) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(exception.PanicLogging)
	}
	userService.UserRepository.Create(model.Username, string(hashPassword), "user")
}
