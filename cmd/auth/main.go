package main

import (
	"database/sql"
	"fmt"
	grpcAuth "github.com/drakenchef/Tinder/internal/pkg/auth/delivery/grpc"
	generatedAuth "github.com/drakenchef/Tinder/internal/pkg/auth/delivery/grpc/gen"
	authRepo "github.com/drakenchef/Tinder/internal/pkg/auth/repo"
	authUsecase "github.com/drakenchef/Tinder/internal/pkg/auth/usecase"
	middleware "github.com/drakenchef/Tinder/internal/pkg/middleware/metrics"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
func run() error {

	logger, err := utils.FileLogger("/var/log/auth_app.log")
	if err != nil {
		return err
	}

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Print(err)
		}
	}(logger)

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
	defer db.Close()

	authRepo := authRepo.NewAuthRepo(db, sugar)
	authUsecase := authUsecase.NewAuthUsecase(authRepo, sugar)
	authHandler := grpcAuth.NewGrpcAuthHandler(authUsecase, sugar)

	srv, ok := net.Listen("tcp", ":8010")
	if ok != nil {
		log.Fatalln("can't listen port", err)
	}

	metricsMw := middleware.NewMetricsMiddleware()
	metricsMw.Register(middleware.ServiceAuthName)

	server := grpc.NewServer(grpc.UnaryInterceptor(metricsMw.ServerMetricsInterceptor))

	generatedAuth.RegisterAuthServer(server, authHandler)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.PathPrefix("/metrics").Handler(promhttp.Handler())

	http.Handle("/", r)
	httpSrv := http.Server{Handler: r, Addr: ":8011"}

	go func() {
		err := httpSrv.ListenAndServe()
		if err != nil {
			fmt.Print(err)
		}
	}()

	fmt.Print("auth running on: ", srv.Addr())
	return server.Serve(srv)
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
