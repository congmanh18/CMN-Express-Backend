package usecase

import (
	entity "cmn-express/domain/user/entity"
	repo "cmn-express/domain/user/repository"
	"context"

	"gorm.io/gorm"
)

type CreateUserUsecase interface {
	ExecuteCreateUser(ctx context.Context, user entity.User) error
}

type CreateUserUsecaseImpl struct {
	userRepo repo.UserRepo
}

func NewCreateUserUsecase(db *gorm.DB) CreateUserUsecase {
	return &CreateUserUsecaseImpl{
		userRepo: repo.NewUserRepo(db),
	}
}

func (c CreateUserUsecaseImpl) ExecuteCreateUser(ctx context.Context, user entity.User) error {
	if err := user.IsValidUser(); err != nil {
		return err
	}
	return c.userRepo.Save(ctx, user)
}
