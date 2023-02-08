package token

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTMaker struct {
}

func NewJWTMaker() (TokenMaker, error) {
	return &JWTMaker{}, nil
}

func (maker *JWTMaker) CreateToken(secretKey string, userID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", payload, err
	}

	JWTToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := JWTToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", payload, err
	}
	return token, payload, nil
}

func (maker *JWTMaker) VerifyToken(secretKey string, token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	}

	JWTToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := JWTToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
