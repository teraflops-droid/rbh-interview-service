package repository

import (
	"context"
	"github.com/teraflops-droid/rbh-interview-service/entity"
)

type UserRepository interface {
	Authentication(ctx context.Context, username string) (entity.User, error)
	Create(ctx context.Context, username string, password string, role string)
}
