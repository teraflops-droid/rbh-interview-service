package service

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/model"
)

type UserService interface {
	Authentication(ctx context.Context, model model.UserModel) entity.User
	Register(model model.UserModel)
}
