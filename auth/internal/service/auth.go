package service

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/repository"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/storage"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

const tokenTTL = 24 * time.Hour

type AuthService interface {
	Register(email, password, name string) (model.User, error)
	Login(email, password string) (model.Token, error)
	ValidateToken(accessToken string) (model.Token, bool, error)
}

type DefaultAuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *DefaultAuthService {
	return &DefaultAuthService{repo: repo}
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

	token := model.Token{
		AccessToken: uuid.NewString(),
		UserID:      user.ID,
		ExpiresAt:   time.Now().UTC().Add(tokenTTL),
	}
	s.repo.StoreToken(token)
	return token, nil
}

func (s *DefaultAuthService) ValidateToken(accessToken string) (model.Token, bool, error) {
	token, err := s.repo.GetToken(accessToken)
	if err != nil {
		return model.Token{}, false, err
	}
	if time.Now().UTC().After(token.ExpiresAt) {
		return token, false, nil
	}
	return token, true, nil
}

func IsNotFound(err error) bool {
	return err == storage.ErrNotFound
}

func IsDuplicate(err error) bool {
	return err == storage.ErrDuplicate
}
