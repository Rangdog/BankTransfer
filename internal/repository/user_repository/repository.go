package repository

import (
	"BankTransfer/internal/model"
	"context"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user model.User) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}
