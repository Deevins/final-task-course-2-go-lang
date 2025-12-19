package service

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/config"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/storage"
	"github.com/gojuno/minimock/v3"
)

func TestAuthService(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "register hashes password",
			run: func(t *testing.T) {
				ctx := context.Background()
				ctrl := minimock.NewController(t)
				t.Cleanup(ctrl.Finish)
				repo := repository.NewAuthRepositoryMock(ctrl)
				jwtConfig := config.JWTConfig{Secret: "test-secret", Expiry: time.Hour}
				service := NewAuthService(repo, jwtConfig)

				repo.CreateUserFunc = func(ctx context.Context, user model.User) (model.User, error) {
					if user.PasswordHash == "" || user.PasswordHash == "password123" {
						t.Fatalf("expected password to be hashed")
					}
					if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte("password123")); err != nil {
						t.Fatalf("expected password hash to match: %v", err)
					}
					return user, nil
				}

				user, err := service.Register(ctx, "user@example.com", "password123", "User")
				if err != nil {
					t.Fatalf("register: %v", err)
				}
				if user.ID == "" {
					t.Fatal("expected user ID to be set")
				}
			},
		},
		{
			name: "login returns token",
			run: func(t *testing.T) {
				ctx := context.Background()
				ctrl := minimock.NewController(t)
				t.Cleanup(ctrl.Finish)
				repo := repository.NewAuthRepositoryMock(ctrl)
				jwtConfig := config.JWTConfig{Secret: "test-secret", Expiry: time.Hour}
				service := NewAuthService(repo, jwtConfig)

				passwordHash := mustHashPassword(t, "password123")
				user := model.User{ID: uuid.NewString(), Email: "user@example.com", PasswordHash: passwordHash}
				repo.GetUserByEmailFn = func(ctx context.Context, email string) (model.User, error) {
					if email != user.Email {
						return model.User{}, storage.ErrNotFound
					}
					return user, nil
				}

				token, err := service.Login(ctx, user.Email, "password123")
				if err != nil {
					t.Fatalf("login: %v", err)
				}
				if token.UserID != user.ID {
					t.Fatalf("expected token user ID %q, got %q", user.ID, token.UserID)
				}
				if token.AccessToken == "" {
					t.Fatal("expected access token to be set")
				}
				if token.ExpiresAt.Before(time.Now().UTC()) {
					t.Fatal("expected token expiry to be in the future")
				}
			},
		},
		{
			name: "login rejects invalid credentials",
			run: func(t *testing.T) {
				ctx := context.Background()
				ctrl := minimock.NewController(t)
				t.Cleanup(ctrl.Finish)
				repo := repository.NewAuthRepositoryMock(ctrl)
				jwtConfig := config.JWTConfig{Secret: "test-secret", Expiry: time.Hour}
				service := NewAuthService(repo, jwtConfig)

				passwordHash := mustHashPassword(t, "password123")
				user := model.User{ID: uuid.NewString(), Email: "user@example.com", PasswordHash: passwordHash}
				repo.GetUserByEmailFn = func(ctx context.Context, email string) (model.User, error) {
					return user, nil
				}

				_, err := service.Login(ctx, user.Email, "wrong")
				if err != ErrInvalidCredentials {
					t.Fatalf("expected invalid credentials error, got %v", err)
				}
			},
		},
		{
			name: "validate token handles expired tokens",
			run: func(t *testing.T) {
				ctx := context.Background()
				ctrl := minimock.NewController(t)
				t.Cleanup(ctrl.Finish)
				repo := repository.NewAuthRepositoryMock(ctrl)
				jwtConfig := config.JWTConfig{Secret: "test-secret", Expiry: time.Hour}
				service := NewAuthService(repo, jwtConfig)

				userID := uuid.NewString()
				expiredToken := makeToken(t, jwtConfig.Secret, userID, time.Now().UTC().Add(-time.Minute))

				token, ok, err := service.ValidateToken(ctx, expiredToken)
				if err != nil {
					t.Fatalf("validate expired token: %v", err)
				}
				if ok {
					t.Fatal("expected expired token to be invalid")
				}
				if token.UserID != userID {
					t.Fatalf("expected expired token user ID %q, got %q", userID, token.UserID)
				}
			},
		},
		{
			name: "validate token accepts valid tokens",
			run: func(t *testing.T) {
				ctx := context.Background()
				ctrl := minimock.NewController(t)
				t.Cleanup(ctrl.Finish)
				repo := repository.NewAuthRepositoryMock(ctrl)
				jwtConfig := config.JWTConfig{Secret: "test-secret", Expiry: time.Hour}
				service := NewAuthService(repo, jwtConfig)

				userID := uuid.NewString()
				validToken := makeToken(t, jwtConfig.Secret, userID, time.Now().UTC().Add(time.Minute))

				token, ok, err := service.ValidateToken(ctx, validToken)
				if err != nil {
					t.Fatalf("validate token: %v", err)
				}
				if !ok {
					t.Fatal("expected token to be valid")
				}
				if token.UserID != userID {
					t.Fatalf("expected token user ID %q, got %q", userID, token.UserID)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}

func makeToken(t *testing.T, secret, subject string, expiresAt time.Time) string {
	t.Helper()
	claims := jwt.RegisteredClaims{
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("sign token: %v", err)
	}
	return signed
}

func mustHashPassword(t *testing.T, value string) string {
	t.Helper()
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}
	return string(hash)
}
