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
	CreateAuthToken(username string, userId uuid.UUID) (token string, err error)
	VerifyAuthToken(token string) (payload *AuthTokenPayload, err error)
}

type AuthTokenPayload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	UserID    uuid.UUID `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p *AuthTokenPayload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return &ErrorTokenExpired{}
	}
	return nil
}

func NewTokenPayload(username string, userID uuid.UUID) (payload *AuthTokenPayload, err error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload = &AuthTokenPayload{
		ID:        tokenID,
		Username:  username,
		UserID:    userID,
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

func (m *TokenMaker) CreateAuthToken(username string, userID uuid.UUID) (token string, err error) {
	payload, err := NewTokenPayload(username, userID)
	if err != nil {
		return "", err
	}
	return m.paseto.Encrypt(m.symmKey, payload, nil)
}

func (m *TokenMaker) VerifyAuthToken(token string) (payload *AuthTokenPayload, err error) {
	payload = &AuthTokenPayload{}

	if err = m.paseto.Decrypt(token, m.symmKey, payload, nil); err != nil {
		return nil, errors.New("service: invalid token")
	}
	if err = payload.Valid(); err != nil {
		return nil, err
	}
	return payload, nil
}
