package model

import "time"

type Account struct {
	Id        int64
	UserId    int64
	Balance   int64
	Currency  string
	CreatedAt time.Time
}
