package token

import (
	"errors"
	"time"

	uuid "github.com/google/uuid"
)

var (
	ErrExpiredTokenDevice = errors.New("token has expired")
	ErrInvalidTokenDevice = errors.New("token is invalid")
)

type PayloadDevice struct {
	ID        uuid.UUID `json:"id"`
	DeviceID  int64     `json:"did"`
	AccountID int64     `json:"aid"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
}

func NewPayloadDevice(deviceId int64, accountId int64, duration time.Duration) (*PayloadDevice, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &PayloadDevice{
		ID:        tokenId,
		DeviceID:  deviceId,
		AccountID: accountId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *PayloadDevice) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
