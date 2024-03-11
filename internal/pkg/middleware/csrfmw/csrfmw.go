package csrfmw

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type HashToken struct {
	Secret []byte
	logger *zap.SugaredLogger
}

func NewHMACKHashToken(secret string, logger *zap.SugaredLogger) (*HashToken, error) {
	return &HashToken{
		Secret: []byte(secret),
		logger: logger,
	}, nil
}

func (tk *HashToken) Create(uid uuid.UUID, tokenExpTime int64) (string, error) {
	tk.logger.Info("Create")
	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%s:%d", uid, tokenExpTime)
	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)
	tk.logger.Info("csrf: ", token)
	tokenData := strings.Split(token, ":")
	tk.logger.Info("Data: ", tokenData, len(tokenData))

	return token, nil
}
func (tk *HashToken) Check(uid uuid.UUID, inputToken string) (bool, error) {
	tk.logger.Info("Check")
	tk.logger.Info("csrf: ", inputToken)
	tokenData := strings.Split(inputToken, ":")
	tk.logger.Info("Data: ", tokenData, len(tokenData))
	if len(tokenData) != 2 {
		return false, errors.New("bad token data")
	}

	tokenExp, err := strconv.ParseInt(tokenData[1], 10, 64)
	if err != nil {
		return false, errors.New("bad token time")
	}
	if tokenExp < time.Now().Unix() {
		return false, errors.New("token expired")
	}
	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%s:%d", uid, tokenExp)
	h.Write([]byte(data))
	expectedMAC := h.Sum(nil)
	messageMAC, err := hex.DecodeString(tokenData[0])
	if err != nil {
		return false, errors.New("can't hex decode string")
	}

	return hmac.Equal(messageMAC, expectedMAC), nil
}
