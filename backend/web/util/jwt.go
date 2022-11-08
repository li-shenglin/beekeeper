package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte("beekeeper.123")
var JwtAuthor = "beekeeper"
var JwtKey = "beekeeper-auth"

type Claims struct {
	jwt.StandardClaims
	UserID   uint   `json:"userId"`
	Nickname string `json:"nickname"`
}

func GenerateToken(userID uint, nickName string) (string, error) {
	expireTime := time.Now().Add(24 * 14 * time.Hour)
	claims := Claims{
		UserID:   userID,
		Nickname: nickName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    JwtAuthor,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok {
			return claims, nil
		}
	}

	return nil, err
}
