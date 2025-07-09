package model

import "time"

type Transfer struct {
	Id            int64
	FromAccountId int64
	ToAccountId   int64
	Amount        int64
	CreateAt      time.Time
}
