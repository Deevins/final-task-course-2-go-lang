package repository

import (
	"context"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"
)

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -i AuthRepository -o auth_repository_minimock.go

type AuthRepository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByID(ctx context.Context, id string) (model.User, error)
}
