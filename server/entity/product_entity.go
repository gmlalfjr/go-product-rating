package entity

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Id          int64
	Title       string
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
	Ratings     []Rating `gorm:"foreignKey:ProductId"`
}
