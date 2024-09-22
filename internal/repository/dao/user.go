package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicate = errors.New("邮箱或手机号冲突")
)

type UserDao interface {
	Insert(ctx context.Context, user User) error
}

type UserGORMDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &UserGORMDao{
		db: db,
	}
}

func (d UserGORMDao) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Ctime = now
	user.Utime = now
	err := d.db.WithContext(ctx).Create(&user).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueConflictsErrNo = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突或者手机号冲突
			return ErrUserDuplicate
		}
	}
	return err
}

type User struct {
	Id       int64          `gorm:"primary_key"`
	Email    sql.NullString `gorm:"unique"`
	Password string
	Phone    sql.NullString `gorm:"unique"`
	Nickname string
	AboutMe  string
	Birthday int64
	Ctime    int64
	Utime    int64
}
