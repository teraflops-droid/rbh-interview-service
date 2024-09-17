package impl

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/teraflops-droid/rbh-interview-service/entity"
	"github.com/teraflops-droid/rbh-interview-service/exception"
	"github.com/teraflops-droid/rbh-interview-service/repository"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) Create(username string, password string, role string) {
	userRole := entity.UserRole{
		Id:       uuid.New(),
		Username: username,
		Role:     role,
	}

	user := entity.User{
		Username:  username,
		Password:  password,
		IsActive:  true,
		UserRoles: userRole,
	}
	err := userRepository.DB.Create(&user).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) DeleteAll() {
	err := userRepository.DB.Where("1=1").Delete(&entity.User{}).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) Authentication(ctx context.Context, username string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).
		Joins("inner join tb_user_role on tb_user_role.username = tb_user.username").
		Preload("UserRoles").
		Where("tb_user.username = ? and tb_user.is_active = ?", username, true).
		Find(&userResult)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return userResult, nil
}
