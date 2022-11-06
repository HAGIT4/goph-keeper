package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func (as *agentService) encrypt(plain []byte) (encypted []byte, err error) {
	c, err := aes.NewCipher(as.encKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plain, nil), nil
}

func (as *agentService) decrypt(encrypted []byte) (plain []byte, err error) {
	c, err := aes.NewCipher(as.encKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return nil, errors.New("service: ciphertext too short")
	}

	nonce, encrypted := encrypted[:nonceSize], encrypted[nonceSize:]
	return gcm.Open(nil, nonce, encrypted, nil)
}
