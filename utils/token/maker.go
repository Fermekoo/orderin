package token

import (
	"time"

	"github.com/google/uuid"
)

type TokenMaker interface {
	CreateToken(secretkey string, userID uuid.UUID, duration time.Duration) (string, *Payload, error)
	VerifyToken(secretKey string, token string) (*Payload, error)
}
