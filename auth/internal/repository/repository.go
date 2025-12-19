package repository

import (
	"context"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByID(ctx context.Context, id string) (model.User, error)
}
