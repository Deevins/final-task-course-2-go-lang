package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/auth/internal/storage"
)

type PostgresAuthRepository struct {
	db *pgxpool.Pool
}

func NewPostgresAuthRepository(db *pgxpool.Pool) *PostgresAuthRepository {
	return &PostgresAuthRepository{db: db}
}

func (r *PostgresAuthRepository) CreateUser(user model.User) (model.User, error) {
	const query = `
		INSERT INTO users (id, email, name, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(
		context.Background(),
		query,
		user.ID,
		user.Email,
		user.Name,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return model.User{}, storage.ErrDuplicate
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *PostgresAuthRepository) GetUserByEmail(email string) (model.User, error) {
	const query = `
		SELECT id, email, name, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1`
	var user model.User
	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, storage.ErrNotFound
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *PostgresAuthRepository) GetUserByID(id string) (model.User, error) {
	const query = `
		SELECT id, email, name, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1`
	var user model.User
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, storage.ErrNotFound
		}
		return model.User{}, err
	}
	return user, nil
}
