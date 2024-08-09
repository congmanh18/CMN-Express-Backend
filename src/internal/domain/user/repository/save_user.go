package repository

import (
	entity "cmn-express/src/internal/domain/user/entity"
	"context"
)

func (u *userRepoImpl) Save(ctx context.Context, user entity.User) error {
	return u.gorm.WithContext(ctx).Debug().Save(&user).Error
}

// // Save implements UserRepo.
// func (u *userRepoImpl) Save(ctx context.Context, user entity.User) error {
// 	panic("unimplemented")
// }
