package main

import (
	"context"
	"database/sql"
	"fmt"
	authHandler "github.com/drakenchef/Tinder/internal/pkg/auth/delivery/http"
	authRepo "github.com/drakenchef/Tinder/internal/pkg/auth/repo"
	authUsecase "github.com/drakenchef/Tinder/internal/pkg/auth/usecase"
	usersHandler "github.com/drakenchef/Tinder/internal/pkg/users/delivery/http"
	usersRepo "github.com/drakenchef/Tinder/internal/pkg/users/repo"
	usersUsecase "github.com/drakenchef/Tinder/internal/pkg/users/usecase"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
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
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	authRepo := authRepo.NewAuthRepo(db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo)
	authHandler := authHandler.NewAuthHandler(authUsecase)

	usersRepo := usersRepo.NewUsersRepo(db)
	usersUsecase := usersUsecase.NewUsersUsecase(usersRepo)
	usersHandler := usersHandler.NewUsersHandler(usersUsecase)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	//r.Use(Check)
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.Handle("/signup", http.HandlerFunc(authHandler.SignUp)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		auth.Handle("/signin", http.HandlerFunc(authHandler.SignIn)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		auth.Handle("/checkauth", http.HandlerFunc(authHandler.CheckAuth)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
	}
	user := r.PathPrefix("/user").Subrouter()
	user.Use(Check)
	{
		user.Handle("/list", http.HandlerFunc(usersHandler.UsersList)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		user.Handle("/profile", http.HandlerFunc(usersHandler.GetUser)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		user.Handle("/update", http.HandlerFunc(usersHandler.UpdateUser)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
		user.Handle("/image", http.HandlerFunc(usersHandler.UpdateUserImage)).
			Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
	}

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	http.Handle("/", r)

	srv := new(Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), r); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("Tinder Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Tinder Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

//func CORSMethodMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//		logrus.Print(req.Host)
//		next.ServeHTTP(w, req)
//	})
//}

func Check(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_, err := utils.CheckAuth(req)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, req)
	})
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// init db

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

//init server

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
