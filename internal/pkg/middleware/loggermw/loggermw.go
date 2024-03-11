package loggermw

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(logger *zap.SugaredLogger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Logging() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			l.logger.Info("req url")
			req.Header.Set("abcd", uuid.New().String())
			next.ServeHTTP(w, req)

		})
	}
}
