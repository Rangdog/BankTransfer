package repository

import (
	user "BankTransfer/internal/repository/user_repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	UserRepository user.UserRepository
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserRepository: user.NewSQLUserRepository(db),
	}
}
