package repo

import (
	entity "cmn-express/domain/user/entity"
	"context"
)

// Find user with user email

func (u *userRepoImpl) Find(ctx context.Context, email string, user entity.User) error {
	return u.gorm.WithContext(ctx).Debug().Where("email = ?", email).First(&user).Error
}
