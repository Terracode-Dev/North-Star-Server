package rba

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenarateJWTkey(exp time.Duration, payload interface{}, sec []byte) (string, error) {
	claim := &jwt.MapClaims{
		"data": payload,
		"exp":  time.Now().Add(exp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	strToken, err := token.SignedString(sec)
	if err != nil {
		return "", err
	}
	return strToken, nil
}

func ValidateJWTkey(token string, sec []byte) (*JWTPayload, error) {
	claims := &JWTPayload{}
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return sec, nil
	})
	if err != nil {
		return nil, err
	}
	if t.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
