package service

import "context"

type UserService interface {
	Sign(ctx context.Context) error
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Sign(ctx context.Context) error {
	return nil
}
