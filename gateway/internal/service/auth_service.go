package service

import (
	"context"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/model"
	authv1 "github.com/Deevins/final-task-course-2-go-lang/gateway/internal/pb/auth/v1"
)

type AuthGatewayService interface {
	SignUp(ctx context.Context, req model.SignUpRequest) (*model.SignUpResponse, error)
	SignIn(ctx context.Context, req model.SignInRequest) (*model.SignInResponse, error)
	ValidateToken(ctx context.Context, accessToken string) (*authv1.ValidateTokenResponse, error)
}

type authGatewayService struct {
	client authv1.AuthServiceClient
}

func NewAuthGatewayService(client authv1.AuthServiceClient) AuthGatewayService {
	if client == nil {
		panic("auth gateway service requires gRPC client")
	}
	return &authGatewayService{client: client}
}

func (s *authGatewayService) SignUp(ctx context.Context, req model.SignUpRequest) (*model.SignUpResponse, error) {
	resp, err := s.client.SignUp(ctx, &authv1.SignUpRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &model.SignUpResponse{UserID: resp.GetUserId()}, nil
}

func (s *authGatewayService) SignIn(ctx context.Context, req model.SignInRequest) (*model.SignInResponse, error) {
	resp, err := s.client.SignIn(ctx, &authv1.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	expiresAt := resp.GetExpiresAt()
	var expiresTime model.SignInResponse
	if expiresAt != nil {
		expiresTime.ExpiresAt = expiresAt.AsTime()
	}
	expiresTime.AccessToken = resp.GetAccessToken()
	return &expiresTime, nil
}

func (s *authGatewayService) ValidateToken(ctx context.Context, accessToken string) (*authv1.ValidateTokenResponse, error) {
	return s.client.ValidateToken(ctx, &authv1.ValidateTokenRequest{AccessToken: accessToken})
}
