package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type TokenMakerInterface interface {
	CreateAuthToken(username string) (token string, err error)
	VerifyAuthToken(token string) (payload *AuthTokenPayload, err error)
}

type AuthTokenPayload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewTokenPayload(username string) (payload *AuthTokenPayload, err error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload = &AuthTokenPayload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}
	return payload, nil
}

type TokenMaker struct {
	paseto  *paseto.V2
	symmKey []byte
}

func NewTokenMaker(symmKey string) (maker *TokenMaker, err error) {
	fmt.Println("SymmKey: ", symmKey)
	if len(symmKey) != chacha20poly1305.KeySize {
		return nil, errors.New("service: invalid key size")
	}

	maker = &TokenMaker{
		paseto:  paseto.NewV2(),
		symmKey: []byte(symmKey),
	}
	return maker, nil
}

func (m *TokenMaker) CreateAuthToken(username string) (token string, err error) {
	payload, err := NewTokenPayload(username)
	if err != nil {
		return "", err
	}
	return m.paseto.Encrypt(m.symmKey, payload, nil)
}

func (m *TokenMaker) VerifyAuthToken(token string) (payload *AuthTokenPayload, err error) {
	payload = &AuthTokenPayload{}

	err = m.paseto.Decrypt(token, m.symmKey, payload, nil)
	if err != nil {
		return nil, errors.New("service: nvalid token")
	}

	// err = payload.Valid()
	// if err != nil {
	// 	return nil, err
	// }

	return payload, nil
}
