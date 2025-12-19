package repository

import (
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/storage"
)

type AuthRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(id string) (model.User, error)
	StoreToken(token model.Token)
	GetToken(accessToken string) (model.Token, error)
}

type InMemoryAuthRepository struct {
	store *storage.InMemoryAuthStorage
}

func NewInMemoryAuthRepository(store *storage.InMemoryAuthStorage) *InMemoryAuthRepository {
	return &InMemoryAuthRepository{store: store}
}

func (r *InMemoryAuthRepository) CreateUser(user model.User) (model.User, error) {
	return r.store.CreateUser(user)
}

func (r *InMemoryAuthRepository) GetUserByEmail(email string) (model.User, error) {
	return r.store.GetUserByEmail(email)
}

func (r *InMemoryAuthRepository) GetUserByID(id string) (model.User, error) {
	return r.store.GetUserByID(id)
}

func (r *InMemoryAuthRepository) StoreToken(token model.Token) {
	r.store.StoreToken(token)
}

func (r *InMemoryAuthRepository) GetToken(accessToken string) (model.Token, error) {
	return r.store.GetToken(accessToken)
}
