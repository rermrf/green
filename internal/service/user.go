package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"green/internal/domain"
	"green/internal/repository"
)

var (
	ErrUserDuplicate = repository.ErrUserDuplicate
)

type UserService interface {
	Signup(ctx context.Context, user domain.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Signup(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return s.repo.Create(ctx, user)
}
