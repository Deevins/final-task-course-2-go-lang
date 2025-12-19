package storage

import (
	"errors"
	"sync"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrDuplicate = errors.New("duplicate")
)

type InMemoryAuthStorage struct {
	mu           sync.RWMutex
	usersByID    map[string]model.User
	usersByEmail map[string]string
}

func NewInMemoryAuthStorage() *InMemoryAuthStorage {
	return &InMemoryAuthStorage{
		usersByID:    make(map[string]model.User),
		usersByEmail: make(map[string]string),
	}
}

func (s *InMemoryAuthStorage) CreateUser(user model.User) (model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.usersByEmail[user.Email]; exists {
		return model.User{}, ErrDuplicate
	}
	s.usersByID[user.ID] = user
	s.usersByEmail[user.Email] = user.ID
	return user, nil
}

func (s *InMemoryAuthStorage) GetUserByEmail(email string) (model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	id, ok := s.usersByEmail[email]
	if !ok {
		return model.User{}, ErrNotFound
	}
	user, ok := s.usersByID[id]
	if !ok {
		return model.User{}, ErrNotFound
	}
	return user, nil
}

func (s *InMemoryAuthStorage) GetUserByID(id string) (model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.usersByID[id]
	if !ok {
		return model.User{}, ErrNotFound
	}
	return user, nil
}
