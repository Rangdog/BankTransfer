package model

import "time"

type User struct {
	Id           int64
	Username     string
	PasswordHash string
	CreateAt     time.Time
}
