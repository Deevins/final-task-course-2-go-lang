package model

import "time"

type User struct {
	ID           string
	Email        string
	Name         string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Token struct {
	AccessToken string
	UserID      string
	ExpiresAt   time.Time
}
