package repository

import (
	"BankTransfer/internal/model"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLUserRepository struct {
	db *pgxpool.Pool
}

func NewSQLUserRepository(db *pgxpool.Pool) UserRepository {
	return &SQLUserRepository{}
}

func (r *SQLUserRepository) RegisterUser(ctx context.Context, user model.User) (*model.User, error) {
	user.CreateAt = time.Now()
	err := r.db.QueryRow(ctx, `INSERT INTO users (username, password_hash, create_at) VALUES ($1,$2,$3) RETURNING id`, user.Username, user.PasswordHash, user.CreateAt).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SQLUserRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	err := r.db.QueryRow(ctx, `select id, password_hash from users where username = $1`, username).Scan(&user.Id, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
