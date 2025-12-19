package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/config"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/storage"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService interface {
	Register(email, password, name string) (model.User, error)
	Login(email, password string) (model.Token, error)
	ValidateToken(accessToken string) (model.Token, bool, error)
}

type DefaultAuthService struct {
	repo      repository.AuthRepository
	jwtConfig config.JWTConfig
}

func NewAuthService(repo repository.AuthRepository, jwtConfig config.JWTConfig) *DefaultAuthService {
	return &DefaultAuthService{repo: repo, jwtConfig: jwtConfig}
}

func (s *DefaultAuthService) Register(email, password, name string) (model.User, error) {
	user := model.User{
		ID:        uuid.NewString(),
		Email:     email,
		Name:      name,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}
	return s.repo.CreateUser(user)
}

func (s *DefaultAuthService) Login(email, password string) (model.Token, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return model.Token{}, err
	}
	if user.Password != password {
		return model.Token{}, ErrInvalidCredentials
	}

	expiresAt := time.Now().UTC().Add(s.jwtConfig.Expiry)
	claims := jwt.RegisteredClaims{
		Subject:   user.ID,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := jwtToken.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken: accessToken,
		UserID:      user.ID,
		ExpiresAt:   expiresAt,
	}, nil
}

func (s *DefaultAuthService) ValidateToken(accessToken string) (model.Token, bool, error) {
	claims := &jwt.RegisteredClaims{}
	parsedToken, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
		}
		return []byte(s.jwtConfig.Secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			expiresAt := time.Time{}
			if claims.ExpiresAt != nil {
				expiresAt = claims.ExpiresAt.Time
			}
			return model.Token{
				AccessToken: accessToken,
				UserID:      claims.Subject,
				ExpiresAt:   expiresAt,
			}, false, nil
		}
		return model.Token{}, false, err
	}
	if !parsedToken.Valid {
		return model.Token{}, false, nil
	}

	expiresAt := time.Time{}
	if claims.ExpiresAt != nil {
		expiresAt = claims.ExpiresAt.Time
	}

	return model.Token{
		AccessToken: accessToken,
		UserID:      claims.Subject,
		ExpiresAt:   expiresAt,
	}, true, nil
}

func IsNotFound(err error) bool {
	return err == storage.ErrNotFound
}

func IsDuplicate(err error) bool {
	return err == storage.ErrDuplicate
}
