package core

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	AuthorizationString string = "Authorization"
	AuthorizationBearer string = "Bearer"
)

type AccessToken struct {
	TokenType             string `json:"token_type"`
	AccessToken           string `json:"access_token"`
	ExpiresAt             int64  `json:"expires_at"`
	IssuedAt              int64  `json:"issued_at"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
	RefreshTokenIssuedAt  int64  `json:"refresh_token_issued_at"`
}

type Claims struct {
	UserId    int
	Username  string
	Refreshed bool
	jwt.StandardClaims
}

func CreateAccessToken(userId int, username string, refresh bool, expire time.Duration) (string, int64, int64, error) {
	now := time.Now()
	expiresAt := now.Add(expire).Unix()
	issuedAt := now.Unix()
	claims := Claims{
		UserId:    userId,
		Username:  username,
		Refreshed: refresh,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Settings.SecretKey))
	if err != nil {
		return "", 0, 0, err
	}
	return tokenString, expiresAt, issuedAt, nil
}
