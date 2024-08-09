package common

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"

	model "cmn-express/src/apis/token/model"
)

var (
	TimeLifeAccessTokenShort = time.Minute * 5
	TimeLifeAccessToken      = time.Second * 30
	TimeLifeRefreshToken     = time.Hour * 24 * 7
)

func GenToken(ID, accountType string, shouldGenRefreshToken bool) (token *model.Token, err error) {
	claimsAccessToken := &model.JwtCustomClaims{
		ID:          ID,
		AccountType: accountType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TimeLifeAccessToken)),
		},
	}

	claimsRefreshToken := &model.JwtCustomClaims{
		ID:          ID,
		AccountType: accountType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TimeLifeRefreshToken)),
		},
	}

	secretKey := []byte(viper.Get("JWT_SECRET_KEY").(string))
	var accessToken, refreshToken *string

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccessToken)
	accessTokenString, err := accessTokenObj.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	accessToken = &accessTokenString

	if shouldGenRefreshToken {
		refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
		refreshTokenString, err := refreshTokenObj.SignedString(secretKey)
		if err != nil {
			return nil, err
		}
		refreshToken = &refreshTokenString
	}

	return &model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GenShortToken(ID, accountType string) (*string, error) {
	claimsShortAccessToken := &model.JwtCustomClaims{
		ID:          ID,
		AccountType: accountType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TimeLifeAccessTokenShort)),
		},
	}
	secretKey := []byte(viper.Get("JWT_SECRET_KEY").(string))
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsShortAccessToken)
	accessTokenString, err := accessTokenObj.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &accessTokenString, nil
}
