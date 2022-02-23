package entity

import (
	"gorm.io/gorm"
	"time"
)

type Rating struct {
	gorm.Model
	Id        int64
	Star      int
	Comment   string
	CreatedAt time.Time
	ProductId uint
}
