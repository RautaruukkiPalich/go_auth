package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rautaruukkipalich/go_auth/config"
	"github.com/rautaruukkipalich/go_auth/internal/model"
)

var (
	jwtSecretKey = []byte(config.JWT_SECRET_KEY)
 	ttlSeconds = config.JWT_TTL_SECONDS
)

func EncodeJWTToken(u *model.User) (string, error){
	payload := jwt.MapClaims{
        "sub":  u.Id,
        "exp":  time.Now().Add(time.Second * time.Duration(ttlSeconds)).Unix(),
    }

    // Создаем новый JWT-токен и подписываем его по алгоритму HS256
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

    signedToken, err := token.SignedString(jwtSecretKey)
    if err != nil {
        return "", errors.New("JWT token failed to signed")
    }
	return signedToken, nil
}

func DecodeJWTToken(token string) (int, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		token,
		claims, 
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecretKey, nil
		},
	)

	if err != nil {
		return 0, errors.New(err.Error())
	}

	sub := int(claims["sub"].(float64))

	return sub, nil
}