package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type jwtConfig struct {
	Secret string
}

var cfg = jwtConfig{
	Secret: "CyanPigeon",
}

type JWTClaim struct {
	UUID int64 `json:"uuid"`
	jwt.RegisteredClaims
}

func GenToken(uuid int64) (string, error) {
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaim{
		UUID: uuid,
	})
	return j.SignedString([]byte(cfg.Secret))
}

func ValidateToken(token string) (bool, JWTClaim, error) {
	j := &JWTClaim{}
	t, err := jwt.ParseWithClaims(token, j, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return false, JWTClaim{}, err
	}
	if t.Valid {
		return true, *j, nil
	}
	return false, JWTClaim{}, errors.New("invalid token")
}
