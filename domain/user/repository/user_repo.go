package repo

import (
	entity "cmn-express/domain/user/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(ctx context.Context, user entity.User) error
	Update(ctx context.Context, id string, user entity.User) error
}

type userRepoImpl struct {
	gorm *gorm.DB
}

func NewUserRepo(gorm *gorm.DB) UserRepo {
	return &userRepoImpl{
		gorm: gorm,
	}
}
