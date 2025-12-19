package model

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time
}

type Token struct {
	AccessToken string
	UserID      string
	ExpiresAt   time.Time
}
