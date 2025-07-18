package token

import "time"

type Maker interface {
	CreateToken(userId int64, accountId int64, accountCode string, permissions []string, modules []string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)

	CreateTokenDevice(deviceId int64, accountId int64, duration time.Duration) (string, *PayloadDevice, error)

	VerifyTokenDevice(token string) (*PayloadDevice, error)
}
