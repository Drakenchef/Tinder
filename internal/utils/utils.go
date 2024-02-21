package utils

import (
	"crypto/rand"
	"encoding/hex"
)

const saltLength = 16

func GenerateRandomSalt() string {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(salt)
}
