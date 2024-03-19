package corsmw

import (
	"github.com/drakenchef/Tinder/internal/pkg/middleware/csrfmw"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type CorsMw struct {
	logger *zap.SugaredLogger
}

func NewCorsMw(logger *zap.SugaredLogger) *CorsMw {
	return &CorsMw{logger: logger}
}

func (cm *CorsMw) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type,X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization,X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_HOST_NAME"))
		//w.Header().Set("Access-Control-Max-Age", "86400")
		if r.Method == http.MethodOptions {
			uidFromContext := r.Header.Get("uid")
			uid, err := uuid.Parse(uidFromContext)
			if err != nil {
				cm.logger.Info("uid not found")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			tokenExpTime := int64(3600)
			csrf, _ := csrfmw.NewHMACKHashToken("zxczxczczxc", cm.logger)
			csrfToken, err := csrf.Create(uid, tokenExpTime)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("X-CSRF-Token", csrfToken)
			return
		}
		next.ServeHTTP(w, r)
	})
}
