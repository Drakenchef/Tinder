package grpc

import (
	"context"

	"go.uber.org/zap"

	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/auth"
	gen "github.com/drakenchef/Tinder/internal/pkg/auth/delivery/grpc/gen"
)

type GrpcAuthHandler struct {
	authUsecase auth.AuthUsecase
	logger      *zap.SugaredLogger

	gen.UnimplementedAuthServer
}

func NewGrpcAuthHandler(authUsecase auth.AuthUsecase, logger *zap.SugaredLogger) *GrpcAuthHandler {
	return &GrpcAuthHandler{
		authUsecase: authUsecase,
		logger:      logger,
	}
}

func (h *GrpcAuthHandler) SignUp(ctx context.Context, req *gen.SignUpRequest) (*gen.SignUpResponse, error) {
	user := models.SignInInput{ // Приведение типов может потребовать изменений в зависимости от вашей модели
		Login:    req.Login,
		Password: req.Password,
	}
	err := h.authUsecase.CreateUser(ctx, user)
	if err != nil {
		h.logger.Info(err)
		return nil, err // Здесь вы можете возвращать кастомные gRPC ошибки используя status.Errorf()
	}

	return &gen.SignUpResponse{}, nil
}

func (h *GrpcAuthHandler) SignIn(ctx context.Context, req *gen.SignInRequest) (*gen.SignInResponse, error) {
	user := models.SignInInput{
		Login:    req.Login,
		Password: req.Password,
	}

	token, err := h.authUsecase.GenerateToken(ctx, user)
	if err != nil {
		h.logger.Info(err)
		return nil, err
	}

	// Пример возможного возвращаемого сообщения
	// Вы можете адаптировать возвращаемую структуру в соответствии с вашим `auth.proto`
	response := &gen.SignInResponse{
		Token: token, // Пример, вы можете иметь другие поля в ответе
	}

	return response, nil
}
