package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomSalt() string {
	saltLength := 16
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(salt)
}
