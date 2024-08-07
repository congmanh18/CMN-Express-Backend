package user

import (
	entity "cmn-express/src/domain/user/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(ctx context.Context, user entity.User) error
}

type userRepoImpl struct {
	gorm *gorm.DB
}

func NewUserRepo(gorm *gorm.DB) UserRepo {
	return &userRepoImpl{
		gorm: gorm,
	}
}
