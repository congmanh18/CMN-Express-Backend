package repo

import (
	entity "cmn-express/domain/user/entity"
	"context"
)

func (u *userRepoImpl) Save(ctx context.Context, user entity.User) error {
	return u.gorm.WithContext(ctx).Debug().Save(&user).Error
}
