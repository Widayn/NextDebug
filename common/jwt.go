package common

import (
	"github.com/golang-jwt/jwt"
	"prmlk.com/nextdebug/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: int64(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Debug.pink",
			Subject:   "user token",
		},
	}

	//priKey, err := jwt.ParseECPrivateKeyFromPEM(jwtKey)
	//if err != nil {
	//	return "", err
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
