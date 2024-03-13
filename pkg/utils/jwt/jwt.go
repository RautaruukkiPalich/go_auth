package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rautaruukkipalich/go_auth/pkg/utils/env"

)

type JWTConfig struct {
	JWT_SECRET_KEY []byte
	JWT_TTL_SECONDS int
}

func newJwtConfig() *JWTConfig {
	return &JWTConfig{
		JWT_SECRET_KEY: []byte(env.GetEnv("JWT_SECRET_KEY", "secretkey")),
		JWT_TTL_SECONDS: env.GetEnvAsInt("JWT_TTL_SECONDS", 3600),
	}
}


var ErrJWTDecode = errors.New("unexpected signing method")
var ErrJWTEncode = errors.New("JWT token failed to signed")


func EncodeJWTToken(id int) (string, error){
	jwtCfg := newJwtConfig()

	payload := jwt.MapClaims{
        "sub":  id,
        "exp":  time.Now().Add(
			time.Second * time.Duration(jwtCfg.JWT_TTL_SECONDS),
			).Unix(),
    }

    // Создаем новый JWT-токен и подписываем его по алгоритму HS256
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

    signedToken, err := token.SignedString(jwtCfg.JWT_SECRET_KEY)
    if err != nil {
        return "", ErrJWTEncode
    }
	return signedToken, nil
}

func DecodeJWTToken(token string) (int, error) {
	jwtCfg := newJwtConfig()

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		token,
		claims, 
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrJWTDecode
			}
			return jwtCfg.JWT_SECRET_KEY, nil
		},
	)

	if err != nil {
		return 0, errors.New(err.Error())
	}

	sub := int(claims["sub"].(float64))

	return sub, nil
}