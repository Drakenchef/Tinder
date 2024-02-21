package auth

import (
	"crypto/md5"
	"encoding/hex"
)

type Auther struct {
}

func (a *Auther) GeneratePasswordHash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}
