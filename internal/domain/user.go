package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string
	Nickname string
	Phone    string
	AboutMe  string
	Birthday time.Time
	Ctime    time.Time
}
