package random

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

func Encode(u uuid.UUID) string {
	return base58.Encode(u[:])
}

func Decode(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(s))
}

func GenerateUUID() string {
	return shortuuid.New()
}
