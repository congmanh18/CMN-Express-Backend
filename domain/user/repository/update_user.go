package repo

import (
	entity "cmn-express/domain/user/entity"
	"context"
)

func (u *userRepoImpl) Update(ctx context.Context, id string, user entity.User) error {
	return u.gorm.WithContext(ctx).Debug().Where("id = ?", id).Updates(&user).Error
}
