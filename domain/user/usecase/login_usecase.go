package usecase

import (
	entity "cmn-express/domain/user/entity"
	repo "cmn-express/domain/user/repository"
	"context"

	"gorm.io/gorm"
)

type LoginUserUsecase interface {
	ExecuteLoginUser(ctx context.Context, email string, user entity.User) error
}

type LoginUserUsecaseImpl struct {
	userRepo repo.UserRepo
}

func NewLoginUserUsecase(db *gorm.DB) LoginUserUsecase {
	return &LoginUserUsecaseImpl{
		userRepo: repo.NewUserRepo(db),
	}
}

func (l LoginUserUsecaseImpl) ExecuteLoginUser(ctx context.Context, email string, user entity.User) error {
	return l.userRepo.Find(ctx, email, user)
}
