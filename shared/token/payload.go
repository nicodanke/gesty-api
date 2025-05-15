package token

import (
	"errors"
	"time"

	uuid "github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID          uuid.UUID `json:"id"`
	UserID      int64     `json:"userId"`
	AccountID   int64     `json:"accountId"`
	AccountCode string    `json:"accountCode"`
	IssuedAt    time.Time `json:"issuedAt"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

func NewPayload(userId int64, accountId int64, accountCode string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:          tokenId,
		UserID:      userId,
		AccountID:   accountId,
		AccountCode: accountCode,
		IssuedAt:    time.Now(),
		ExpiredAt:   time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
