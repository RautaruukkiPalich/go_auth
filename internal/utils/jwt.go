package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rautaruukkipalich/go_auth/internal/model"
)

type JwtConfig struct {
	JWT_SECRET_KEY []byte
	JWT_TTL_SECONDS int
}

func newJwtConfig() *JwtConfig {
	return &JwtConfig{
		JWT_SECRET_KEY: []byte(GetEnv("JWT_SECRET_KEY", "secretkey")),
		JWT_TTL_SECONDS: GetEnvAsInt("JWT_TTL_SECONDS", 3600),
	}
}

func EncodeJWTToken(u *model.User) (string, error){
	cfg := newJwtConfig()
	JWT_SECRET_KEY  := []byte(cfg.JWT_SECRET_KEY)
	JWT_TTL_SECONDS := cfg.JWT_TTL_SECONDS

	payload := jwt.MapClaims{
        "sub":  u.Id,
        "exp":  time.Now().Add(time.Second * time.Duration(JWT_TTL_SECONDS)).Unix(),
    }

    // Создаем новый JWT-токен и подписываем его по алгоритму HS256
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

    signedToken, err := token.SignedString(JWT_SECRET_KEY)
    if err != nil {
        return "", errors.New("JWT token failed to signed")
    }
	return signedToken, nil
}

func DecodeJWTToken(token string) (int, error) {
	cfg := newJwtConfig()
	JWT_SECRET_KEY  := []byte(cfg.JWT_SECRET_KEY)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		token,
		claims, 
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return JWT_SECRET_KEY, nil
		},
	)

	if err != nil {
		return 0, errors.New(err.Error())
	}

	sub := int(claims["sub"].(float64))

	return sub, nil
}