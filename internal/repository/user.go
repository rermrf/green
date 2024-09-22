package repository

import (
	"context"
	"database/sql"
	"green/internal/domain"
	"green/internal/repository/dao"
)

var (
	ErrUserDuplicate = dao.ErrUserDuplicate
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
}

type CachedUserRepository struct {
	dao dao.UserDao
}

func (r *CachedUserRepository) Create(ctx context.Context, user domain.User) error {
	return r.dao.Insert(ctx, r.toEntity(user))
}

func NewCachedUserRepository(dao dao.UserDao) UserRepository {
	return &CachedUserRepository{
		dao: dao,
	}
}

func (r *CachedUserRepository) toEntity(user domain.User) dao.User {
	return dao.User{
		Id: user.Id,
		Email: sql.NullString{
			String: user.Email,
			Valid:  user.Email != "",
		},
		Password: user.Password,
		Phone: sql.NullString{
			String: user.Phone,
			Valid:  user.Phone != "",
		},
		Nickname: user.Nickname,
		AboutMe:  user.AboutMe,
		Birthday: user.Birthday.UnixMilli(),
	}
}
