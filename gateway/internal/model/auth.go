package model

import "time"

// SignUpRequest описывает запрос на регистрацию пользователя.
type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"secret"`
	Name     string `json:"name" binding:"required" example:"Иван Иванов"`
}

// SignUpResponse описывает ответ после регистрации.
type SignUpResponse struct {
	UserID string `json:"user_id" example:"11111111-1111-1111-1111-111111111111"`
}

// SignInRequest описывает запрос на вход.
type SignInRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"secret"`
}

// SignInResponse описывает ответ после входа.
type SignInResponse struct {
	AccessToken string    `json:"access_token" example:"eyJhbGciOi..."`
	ExpiresAt   time.Time `json:"expires_at" example:"2024-01-01T10:00:00Z"`
}
