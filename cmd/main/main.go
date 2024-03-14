package main

import (
	"context"
	"database/sql"
	"fmt"
	authHandler "github.com/drakenchef/Tinder/internal/pkg/auth/delivery/http"
	authRepo "github.com/drakenchef/Tinder/internal/pkg/auth/repo"
	authUsecase "github.com/drakenchef/Tinder/internal/pkg/auth/usecase"
	likesHandler "github.com/drakenchef/Tinder/internal/pkg/likes/delivery/http"
	likesRepo "github.com/drakenchef/Tinder/internal/pkg/likes/repo"
	likesUsecase "github.com/drakenchef/Tinder/internal/pkg/likes/usecase"
	"github.com/drakenchef/Tinder/internal/pkg/middleware/corsmw"
	"github.com/drakenchef/Tinder/internal/pkg/middleware/cspxssmw"
	csrfToken "github.com/drakenchef/Tinder/internal/pkg/middleware/csrfmw"
	"github.com/drakenchef/Tinder/internal/pkg/middleware/loggermw"
	usersHandler "github.com/drakenchef/Tinder/internal/pkg/users/delivery/http"
	usersRepo "github.com/drakenchef/Tinder/internal/pkg/users/repo"
	usersUsecase "github.com/drakenchef/Tinder/internal/pkg/users/usecase"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := InitConfig(); err != nil {
		sugar.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		sugar.Fatalf("error occured while running http server: %s", err.Error())
	}
	db, err := NewPostgresDB(Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		sugar.Fatalf("failed to initialize db: %s", err.Error())
	}
	authRepo := authRepo.NewAuthRepo(db, sugar)
	authUsecase := authUsecase.NewAuthUsecase(authRepo, sugar)
	authHandler := authHandler.NewAuthHandler(authUsecase, sugar)

	usersRepo := usersRepo.NewUsersRepo(db, sugar)
	usersUsecase := usersUsecase.NewUsersUsecase(usersRepo, sugar)
	usersHandler := usersHandler.NewUsersHandler(usersUsecase, sugar)

	likesRepo := likesRepo.NewLikesRepo(db, sugar)
	likesUsecase := likesUsecase.NewLikesUsecase(likesRepo, sugar)
	likesHandler := likesHandler.NewLikesHandler(likesUsecase, sugar)

	cspXssMw := cspxssmw.NewCspXssMW()
	hmackHashToken, _ := csrfToken.NewHMACKHashToken("zxczxczczxc", sugar)
	mylogger := loggermw.NewLogger(sugar)
	corsmw := corsmw.NewCorsMw()

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(mylogger.Logging())
	r.Use(cspXssMw.MiddlewareCSP)
	r.Use(cspXssMw.MiddlewareXSS)
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.Handle("/signup", http.HandlerFunc(authHandler.SignUp)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		auth.Handle("/signin", http.HandlerFunc(authHandler.SignIn)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)

	}
	checkauth := r.PathPrefix("/checkauth").Subrouter()
	checkauth.Use(Check)
	{
		checkauth.Handle("/checkauth", http.HandlerFunc(authHandler.CheckAuth)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
	}
	user := r.PathPrefix("/user").Subrouter()
	user.Use(corsmw.CorsMiddleware)
	user.Use(MiddlewareCSRFCheck(hmackHashToken, sugar))

	{
		user.Handle("/list", http.HandlerFunc(usersHandler.UsersList)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		user.Handle("/profile", http.HandlerFunc(usersHandler.GetUser)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		user.Handle("/update", http.HandlerFunc(usersHandler.UpdateUser)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		user.Handle("/image", http.HandlerFunc(usersHandler.UpdateUserImage)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		user.Handle("/password", http.HandlerFunc(usersHandler.UpdateUserPassword)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		user.Handle("/deleteuser", http.HandlerFunc(usersHandler.DeleteUser)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
	}
	likes := r.PathPrefix("/likes").Subrouter()
	likes.Use(Check)
	{
		likes.Handle("/like", http.HandlerFunc(likesHandler.LikeUser)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		likes.Handle("/mutuallike", http.HandlerFunc(likesHandler.MutualLikeUser)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
	}

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	http.Handle("/", r)

	srv := new(Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), r); err != nil {
			sugar.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	sugar.Info("Tinder started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	sugar.Info("Tinder shutdown")

	if err := srv.Shutdown(context.Background()); err != nil {
		sugar.Fatalf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		sugar.Fatalf("error occured on db connection close: %s", err.Error())
	}

}

func MiddlewareCSRFCheck(hmackHashToken *csrfToken.HashToken, logger *zap.SugaredLogger) func(http.Handler) http.Handler {
	logger.Info("CSRF MIDDLEWARE STARTED")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uid, err := utils.CheckAuth(r)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			r.Header.Set("uid", uid.String())
			if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
				csrfToken := r.Header.Get("X-CSRF-Token")
				logger.Info(csrfToken)
				valid, _ := hmackHashToken.Check(uid, csrfToken)
				if !valid {
					http.Error(w, err.Error(), http.StatusForbidden)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
func MiddlewareCSRFSet(hmackHashToken *csrfToken.HashToken) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenExpTime := int64(3600)
			uid, err := utils.CheckAuth(r)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
				csrfToken, err := hmackHashToken.Create(uid, tokenExpTime)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				w.Header().Set("X-CSRF-Token", csrfToken)
			}
			next.ServeHTTP(w, r)
		})
	}
}
func Check(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		uid, err := utils.CheckAuth(req)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		req.Header.Set("uid", uid.String())
		next.ServeHTTP(w, req)
	})
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
