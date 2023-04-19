package service

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hasher struct {
	constant string
}

func NewHasher(a string) *Hasher {
	return &Hasher{a}
}

func (h *Hasher) Hash(password string) string {
	pass := password + h.constant
	hash := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(hash[:])
}
