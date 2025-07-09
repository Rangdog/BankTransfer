package service

import (
	"BankTransfer/internal/model"
	repository "BankTransfer/internal/repository/user_repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, username string, password string) (*model.User, error) {
	passwordHash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := model.User{
		PasswordHash: passwordHash,
		Username:     username,
	}
	userCreated, err := s.repo.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return userCreated, err
}

func (s *UserService) Login(ctx context.Context, username, password string) error {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if !CheckPassword(user.PasswordHash, password) {
		return errors.New("password doesn't match")
	}
	return nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
