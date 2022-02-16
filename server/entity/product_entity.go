package entity

import (
	"time"
)

type Product struct {
	Id int64
	Title string
	Description string
	CreatedAt time.Time
	ModifiedAt time.Time
}


