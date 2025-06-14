package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"github.com/labstack/gommon/log"
)

func GenerateRandomToken() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Warnf("Failed to generate random bytes")
	}

	hash := sha256.Sum256(randomBytes)

	return hex.EncodeToString(hash[:])
}
