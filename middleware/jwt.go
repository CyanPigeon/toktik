package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTConfig struct {
	Secret              string
	TokenExpireDuration time.Duration
}

type JWTClaim struct {
	UUID string
	jwt.RegisteredClaims
}

func GenToken(cfg JWTConfig, uuid string) (string, error) {
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaim{
		UUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.TokenExpireDuration)),
		}})
	return j.SignedString([]byte(cfg.Secret))
}

func ValidateToken(cfg JWTConfig, token string) (bool, error) {
	var j = new(JWTClaim)
	t, err := jwt.ParseWithClaims(token, j, func(token *jwt.Token) (i interface{}, err error) {
		return cfg.Secret, nil
	})
	if err != nil {
		return false, err
	}
	if t.Valid {
		return true, nil
	}
	return false, errors.New("invalid token")
}
