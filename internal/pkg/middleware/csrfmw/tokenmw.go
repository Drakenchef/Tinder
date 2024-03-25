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
	"sync"
	"time"
)

type HashToken struct {
	Secret   []byte
	logger   *zap.SugaredLogger
	tokens   map[string]struct{} // Сет для хранения активных токенов
	tokensMu sync.RWMutex
}

func NewHMACKHashToken(secret string, logger *zap.SugaredLogger) (*HashToken, error) {
	return &HashToken{
		Secret: []byte(secret),
		logger: logger,
		tokens: make(map[string]struct{}), // инициализируем мапу
	}, nil
}

func (tk *HashToken) Create(uid uuid.UUID, tokenExpTime int64) (string, error) {
	tk.logger.Info("Create")
	h := hmac.New(sha256.New, []byte(tk.Secret))
	currentUnixTime := time.Now().Unix()
	absoluteTokenExpTime := currentUnixTime + tokenExpTime
	data := fmt.Sprintf("%s:%d", uid, absoluteTokenExpTime)
	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(absoluteTokenExpTime, 10)
	tk.logger.Info("csrf: ", token)
	tokenData := strings.Split(token, ":")
	tk.logger.Info("Data: ", tokenData, len(tokenData))

	tk.tokensMu.Lock()
	tk.tokens[token] = struct{}{}
	tk.tokensMu.Unlock()
	tk.logger.Info("TOKENS: ", tk.tokens)
	return token, nil
}
func (tk *HashToken) Check(uid uuid.UUID, inputToken string) (bool, error) {
	tk.logger.Info("Check")

	//tk.tokensMu.RLock()
	//if _, exists := tk.tokens[inputToken]; !exists {
	//	tk.tokensMu.RUnlock()
	//	return false, errors.New("token does not exist")
	//}
	//tk.tokensMu.RUnlock()

	tk.tokensMu.Lock()
	defer tk.tokensMu.Unlock()
	_, found := tk.tokens[inputToken]
	if found {
		delete(tk.tokens, inputToken) // Удаление токена, если он должен быть одноразовым
	}
	tk.logger.Info("TOKENS: ", tk.tokens)
	tk.logger.Info("csrf: ", inputToken)
	if len(inputToken) == 0 {
		return false, errors.New("0 len token")
	}
	tokenData := strings.Split(inputToken, ":")
	tk.logger.Info("Data: ", tokenData, len(tokenData))
	if len(tokenData) != 2 {
		return false, errors.New("bad token data")
	}
	tk.logger.Info("token data checked")
	tokenExp, err := strconv.ParseInt(tokenData[1], 10, 64)
	if err != nil {
		return false, errors.New("bad token time")
	}
	tk.logger.Info("token time checked")
	if tokenExp < time.Now().Unix() {
		return false, errors.New("token expired")
	}
	tk.logger.Info("token exp checked")
	h := hmac.New(sha256.New, []byte(tk.Secret))
	tk.logger.Info("hmac.new")
	data := fmt.Sprintf("%s:%d", uid, tokenExp)
	h.Write([]byte(data))
	tk.logger.Info("h.write")
	expectedMAC := h.Sum(nil)
	messageMAC, err := hex.DecodeString(tokenData[0])
	if err != nil {
		return false, errors.New("can't hex decode string")
	}
	tk.logger.Info("END OF Check", messageMAC, expectedMAC)

	//tk.tokensMu.Lock()
	//delete(tk.tokens, inputToken)
	//tk.tokensMu.Unlock()
	//tk.logger.Info(tk.tokens)

	return true, nil
}
