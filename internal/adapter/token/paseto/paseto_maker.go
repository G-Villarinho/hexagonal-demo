package paseto

import (
	"context"
	"errors"
	"time"

	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type pasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (port.TokenMaker, error) {
	if len(symmetricKey) != 32 {
		return nil, errors.New("chave simétrica deve ter 32 caracteres")
	}

	return &pasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (maker *pasetoMaker) Generate(ctx context.Context, userID string) (string, error) {
	jsonToken := paseto.JSONToken{
		IssuedAt:   time.Now(),
		Expiration: time.Now().Add(24 * time.Hour),
		Subject:    userID,
		Jti:        uuid.NewString(),
		NotBefore:  time.Now(),
	}

	return maker.paseto.Encrypt(maker.symmetricKey, jsonToken, nil)
}

func (maker *pasetoMaker) Verify(ctx context.Context, token string) (string, error) {
	var jsonToken paseto.JSONToken
	err := maker.paseto.Decrypt(token, maker.symmetricKey, &jsonToken, nil)
	if err != nil {
		return "", errors.New("token inválido")
	}

	if err := jsonToken.Validate(); err != nil {
		return "", err
	}

	return jsonToken.Subject, nil
}
