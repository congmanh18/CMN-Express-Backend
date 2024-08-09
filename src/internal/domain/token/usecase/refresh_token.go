package usecase

import (
	"cmn-express/common"
	model "cmn-express/src/apis/token/model"
	"context"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type RefreshTokenUseCase interface {
	ExecRefreshToken(ctx context.Context, token string) (*model.Token, *common.Error)
}

type refreshTokenUseCaseImpl struct {
}

var _ RefreshTokenUseCase = (*refreshTokenUseCaseImpl)(nil)

func NewRefreshTokenUseCase() RefreshTokenUseCase {
	return &refreshTokenUseCaseImpl{}
}

func (r refreshTokenUseCaseImpl) ExecRefreshToken(ctx context.Context, refreshToken string) (*model.Token, *common.Error) {
	parsedRefreshClaims, _ := jwt.ParseWithClaims(refreshToken, &model.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.Get("JWT_SECRET_KEY").(string)), nil
		})

	// refresh token is expired
	if !parsedRefreshClaims.Valid {
		return nil, &common.Error{
			Message:      ErrMsgInvalidToken,
			DebugMessage: "Jwt malformed",
			Code:         ErrCodeInvalidToken,
		}
	}

	claims := parsedRefreshClaims.Claims.(*model.JwtCustomClaims)
	token, err := common.GenToken(claims.ID, claims.AccountType, false)
	if err != nil {
		return nil, &common.Error{
			Message:      ErrMsgGenToken,
			DebugMessage: err.Error(),
			Code:         ErrCodeGenToken,
		}
	}

	return token, nil
}
