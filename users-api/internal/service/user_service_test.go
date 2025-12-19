package service

import (
	"context"
	"testing"
	"time"

	"github.com/Deevins/final-task-course-2-go-lang/users-api/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/users-api/internal/repository"
)

func TestUserService(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "create update delete flow",
			run: func(t *testing.T) {
				ctx := context.Background()
				repo := repository.NewInMemoryUserRepository(nil)
				svc := NewUserService(repo)

				created, err := svc.CreateUser(ctx, model.CreateUserRequest{
					Username: "bob",
					Email:    "bob@example.com",
				})
				if err != nil {
					t.Fatalf("create user: %v", err)
				}
				if created.ID == "" {
					t.Fatal("expected user ID to be set")
				}
				if created.CreatedAt.IsZero() || created.UpdatedAt.IsZero() {
					t.Fatal("expected timestamps to be set")
				}
				if !created.CreatedAt.Equal(created.UpdatedAt) {
					t.Fatal("expected created and updated timestamps to match")
				}

				fetched, err := svc.GetUser(ctx, created.ID)
				if err != nil {
					t.Fatalf("get user: %v", err)
				}
				if fetched.Username != "bob" {
					t.Fatalf("expected username 'bob', got %q", fetched.Username)
				}

				newName := "bobby"
				newEmail := "bobby@example.com"
				updated, err := svc.UpdateUser(ctx, created.ID, model.UpdateUserRequest{
					Username: &newName,
					Email:    &newEmail,
				})
				if err != nil {
					t.Fatalf("update user: %v", err)
				}
				if updated.Username != newName {
					t.Fatalf("expected updated username %q, got %q", newName, updated.Username)
				}
				if updated.Email == nil || *updated.Email != newEmail {
					t.Fatalf("expected updated email %q, got %v", newEmail, updated.Email)
				}
				if !updated.UpdatedAt.After(created.UpdatedAt) {
					t.Fatalf("expected updated timestamp after create, got %s vs %s", updated.UpdatedAt, created.UpdatedAt)
				}

				if err := svc.DeleteUser(ctx, created.ID); err != nil {
					t.Fatalf("delete user: %v", err)
				}

				_, err = svc.GetUser(ctx, created.ID)
				if err == nil {
					t.Fatal("expected error when fetching deleted user")
				}
			},
		},
		{
			name: "list users includes newly created",
			run: func(t *testing.T) {
				ctx := context.Background()
				repo := repository.NewInMemoryUserRepository(map[string]model.User{})
				svc := NewUserService(repo)

				_, err := svc.CreateUser(ctx, model.CreateUserRequest{
					Username: "carol",
					Email:    "carol@example.com",
				})
				if err != nil {
					t.Fatalf("create user: %v", err)
				}

				users, err := svc.ListUsers(ctx)
				if err != nil {
					t.Fatalf("list users: %v", err)
				}
				if len(users) < 2 {
					t.Fatalf("expected at least 2 users (including default), got %d", len(users))
				}

				found := false
				for _, user := range users {
					if user.Username == "carol" {
						found = true
						break
					}
				}
				if !found {
					t.Fatal("expected to find newly created user in list")
				}
			},
		},
		{
			name: "create user sets timestamp bounds",
			run: func(t *testing.T) {
				ctx := context.Background()
				repo := repository.NewInMemoryUserRepository(nil)
				svc := NewUserService(repo)

				before := time.Now().UTC()
				created, err := svc.CreateUser(ctx, model.CreateUserRequest{
					Username: "dave",
					Email:    "dave@example.com",
				})
				if err != nil {
					t.Fatalf("create user: %v", err)
				}
				after := time.Now().UTC()

				if created.CreatedAt.Before(before) || created.CreatedAt.After(after) {
					t.Fatalf("expected created timestamp between %s and %s, got %s", before, after, created.CreatedAt)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}
