package usecase

import (
	entity "cmn-express/domain/user/entity"
	repo "cmn-express/domain/user/repository"
	"context"

	"gorm.io/gorm"
)

type UpdateUserUsecase interface {
	ExecuteUpdateUser(ctx context.Context, id string, user entity.User) error
}

type UpdateUserUsecaseImpl struct {
	userRepo repo.UserRepo
}

func NewUpdateUserUsecase(db *gorm.DB) UpdateUserUsecase {
	return &UpdateUserUsecaseImpl{
		userRepo: repo.NewUserRepo(db),
	}
}

func (u UpdateUserUsecaseImpl) ExecuteUpdateUser(ctx context.Context, id string, user entity.User) error {
	if err := user.IsValidUser(); err != nil {
		return err
	}
	return u.userRepo.Update(ctx, id, user)
}
