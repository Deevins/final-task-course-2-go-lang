package repository

import "github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"

type AuthRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(id string) (model.User, error)
}
