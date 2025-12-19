package grpcserver

import (
	"context"

	pb "github.com/Deevins/final-task-course-2-go-lang/auth/internal/pb/auth/v1"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	authService service.AuthService
}

var _ pb.AuthServiceServer = (*AuthServer)(nil)

func NewAuthServer(svc service.AuthService) *AuthServer {
	return &AuthServer{authService: svc}
}

func (s *AuthServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	user, err := s.authService.Register(ctx, req.GetEmail(), req.GetPassword(), req.GetName())
	if err != nil {
		if service.IsDuplicate(err) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Errorf(codes.Internal, "register: %v", err)
	}

	return &pb.SignUpResponse{UserId: user.ID}, nil
}

func (s *AuthServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	token, err := s.authService.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if err == service.ErrInvalidCredentials {
			return nil, status.Error(codes.Unauthenticated, "invalid credentials")
		}
		if service.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "login: %v", err)
	}

	return &pb.SignInResponse{
		AccessToken: token.AccessToken,
		ExpiresAt:   timestamppb.New(token.ExpiresAt),
	}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	if req.GetAccessToken() == "" {
		return nil, status.Error(codes.InvalidArgument, "access_token is required")
	}

	token, valid, err := s.authService.ValidateToken(ctx, req.GetAccessToken())
	if err != nil {
		if service.IsNotFound(err) {
			return &pb.ValidateTokenResponse{Valid: false}, nil
		}
		return nil, status.Errorf(codes.Internal, "validate token: %v", err)
	}

	return &pb.ValidateTokenResponse{
		Valid:     valid,
		UserId:    token.UserID,
		ExpiresAt: timestamppb.New(token.ExpiresAt),
	}, nil
}
